package ratelimit

import (
	"log"
	"sync/atomic"
	"time"
)

// MaxUint holds the maximum unsigned int value
const MaxUint = ^uint(0)

// MaxInt holds the maximum int value
const MaxInt = int(MaxUint >> 1)

// Manager implements a rate limiter interface.
// The makeToken field is a factory function for creating tokens that will allow different rate limiter implementations
// to define their own custom logic for token creation.
// The makeToken type is tokenFactory, it is a factory function for NewToken.
type Manager struct {
	errorChan    chan error
	ReleaseChan  chan *Token
	outChan      chan *Token
	InChan       chan struct{}
	needToken    int64
	activeTokens map[string]*Token
	limit        int
	makeToken    tokenFactory
}

// Acquire is called to acquire a new token
// When Acquire() is called, an empty struct{} is sent to the in channel,
// and we wait for either a token from the out channel or an error from the error channel.
func (m *Manager) Acquire() (*Token, error) {
	go func() {
		m.InChan <- struct{}{}
	}()

	// Await rate limit token
	select {
	case t := <-m.outChan:
		return t, nil
	case err := <-m.errorChan:
		return nil, err
	}
}

// Release is called to release an active token
func (m *Manager) Release(t *Token) {
	if t.IsExpired() {
		go func() {
			m.ReleaseChan <- t
		}()
	}
}

// NewManager creates a manager type.
func NewManager(conf *Config) *Manager {
	m := &Manager{
		errorChan:    make(chan error),
		outChan:      make(chan *Token),
		InChan:       make(chan struct{}),
		activeTokens: make(map[string]*Token),
		ReleaseChan:  make(chan *Token),
		needToken:    0,
		limit:        conf.Limit,
		makeToken:    NewToken,
	}

	// If limit is not defined, then default to max value
	if m.limit <= 0 {
		m.limit = MaxInt
	}

	// If the config TokenResetsAfter value exists, then run the reset task
	if conf.TokenResetsAfter > 0 {
		m.runResetTokenTask(conf.TokenResetsAfter)
	}

	return m
}

func (m *Manager) incNeedToken() {
	atomic.AddInt64(&m.needToken, 1)
}

func (m *Manager) decNeedToken() {
	atomic.AddInt64(&m.needToken, -1)
}

func (m *Manager) awaitingToken() bool {
	return atomic.LoadInt64(&m.needToken) > 0
}

// TryGenerateToken is used to create a new token and sends it to the out channel.
func (m *Manager) TryGenerateToken() {
	// panic if token factory is not defined
	if m.makeToken == nil {
		panic("ErrTokenFactoryNotDefined")
	}

	// cannot continue if limit has been reached
	if m.isLimitExceeded() {
		m.incNeedToken()
		return
	}

	token := m.makeToken()

	// Add token to active map
	m.activeTokens[token.ID] = token

	// send token to outChan
	go func() {
		m.outChan <- token
	}()
}

func (m *Manager) isLimitExceeded() bool {
	if len(m.activeTokens) >= m.limit {
		return true
	}
	return false
}

func (m *Manager) ReleaseToken(token *Token) {
	if token == nil {
		log.Print("unable to relase nil token")
		return
	}

	if _, ok := m.activeTokens[token.ID]; !ok {
		log.Printf("unable to relase token %s - not in use", token)
		return
	}

	if !token.IsExpired() {
		log.Printf("unable to relase token %s - has not expired", token)
		return
	}

	// Delete from map
	delete(m.activeTokens, token.ID)

	// process anything waiting for a rate limit
	if m.awaitingToken() {
		m.decNeedToken()
		go m.TryGenerateToken()
	}
}

// loops over active tokens and releases any that are expired
func (m *Manager) releaseExpiredTokens() {
	for _, token := range m.activeTokens {
		if token.IsExpired() {
			go func(t *Token) {
				m.ReleaseChan <- t
			}(token)
		}
	}
}

// reset task that runs once per provided duration and releases
// and tokens that need to be reset
func (m *Manager) runResetTokenTask(resetAfter time.Duration) {
	go func() {
		ticker := time.NewTicker(resetAfter)
		for range ticker.C {
			for _, token := range m.activeTokens {
				if token.NeedReset(resetAfter) {
					go func(t *Token) {
						m.ReleaseChan <- t
					}(token)
				}
			}
		}
	}()
}

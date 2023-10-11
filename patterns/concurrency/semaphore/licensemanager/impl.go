package licensemanager

import (
	"errors"
	"sync"
)

type LicenseManager struct {
	sem     chan struct{}
	mtx     sync.Mutex
	maxLic  int
	usedLic int
}

func NewLicenseManager(maxLicenses int) *LicenseManager {
	return &LicenseManager{
		sem:    make(chan struct{}, maxLicenses),
		maxLic: maxLicenses,
	}
}

func (lm *LicenseManager) Acquire() error {
	lm.mtx.Lock()
	defer lm.mtx.Unlock()

	if lm.usedLic >= lm.maxLic {
		return errors.New("all licenses are in use")
	}
	lm.sem <- struct{}{}
	lm.usedLic++
	return nil
}

func (lm *LicenseManager) Release() error {
	lm.mtx.Lock()
	defer lm.mtx.Unlock()

	select {
	case <-lm.sem:
		lm.usedLic--
		return nil
	default:
		return errors.New("no license to release")
	}
}

func (lm *LicenseManager) UsedLicenses() int {
	lm.mtx.Lock()
	defer lm.mtx.Unlock()
	return lm.usedLic
}

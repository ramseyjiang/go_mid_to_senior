package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// State the FSM state for turnstile
type State uint32

const (
	// Locked state
	Locked State = iota
	// Unlocked unlocked state
	Unlocked
)

const (
	// CmdCoin command coin
	CmdCoin = "coin"
	// CmdPush command push
	CmdPush = "push"
)

// Turnstile the finite state machine
type Turnstile struct {
	State State
}

// ExecuteCmd execute command
func (p *Turnstile) ExecuteCmd(cmd string) {
	// get function from transition table
	tuple := CmdStateTuple{strings.TrimSpace(cmd), p.State}
	if f := StateTransitionTable[tuple]; f == nil {
		log.Println("unknown command, try again please")
	} else {
		f(&p.State)
	}
}

func main() {
	machine := &Turnstile{State: Locked}
	prompt(machine.State)
	reader := bufio.NewReader(os.Stdin)

	for {
		// read command from stdin
		cmd, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}

		machine.ExecuteCmd(cmd)
	}
}

// CmdStateTuple tuple for state-command combination
type CmdStateTuple struct {
	Cmd   string
	State State
}

// TransitionFunc transition function
type TransitionFunc func(state *State)

// StateTransitionTable transition table
var StateTransitionTable = map[CmdStateTuple]TransitionFunc{
	{CmdCoin, Locked}: func(state *State) {
		log.Println("unlocked, ready for pass through")
		*state = Unlocked
	},
	{CmdPush, Locked}: func(state *State) {
		log.Println("not allowed, unlock first")
	},
	{CmdCoin, Unlocked}: func(state *State) {
		log.Println("well, don't waste your coin")
	},
	{CmdPush, Unlocked}: func(state *State) {
		log.Println("pass through, shift back to locked")
		*state = Locked
	},
}

func prompt(s State) {
	m := map[State]string{
		Locked:   "Locked",
		Unlocked: "Unlocked",
	}
	log.Printf("current state is [%s], please input command [coin|push]\n", m[s])
}

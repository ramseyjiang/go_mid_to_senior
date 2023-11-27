package cmdevent

import (
	"bufio"
	"io"
	"strings"
)

// step 1
type Command struct {
	Kind string
	Args []string
}

// step 2
type Listener[T any] func(T)

// step 3
type CMDManager[T any] interface {
	Add(n string, l Listener[T]) // Attach a new listener to an event
	Run(inputReader io.Reader)   // Execute the manager.
}

// step 4
type BaseManager[T any] struct {
	lst map[string][]Listener[T]
}

func (m *BaseManager[T]) Invoke(n string, args T) {
	for _, ls := range m.lst[n] {
		ls(args)
	}
}

func (m *BaseManager[T]) Add(n string, l Listener[T]) {
	if m.lst == nil {
		m.lst = make(map[string][]Listener[T])
	}
	m.lst[n] = append(m.lst[n], l)
}

// step 5
type CommandEventManager struct {
	BaseManager[*Command]
}

// Ensure CommandEventManager implements the CMDManager interface
// It can be ignored.
var _ CMDManager[*Command] = &CommandEventManager{}

func (m *CommandEventManager) Run(inputReader io.Reader) {
	scanner := bufio.NewScanner(inputReader)
	for scanner.Scan() {
		input := scanner.Text()
		cmd := strings.Split(input, ":")
		l := len(cmd)

		var args Command
		if l == 0 {
			m.Invoke("no-command", nil)
			continue
		}

		if l > 1 {
			args.Args = strings.Split(cmd[1], " ")
		}

		args.Kind = cmd[0]
		m.Invoke("any-command", &args)
		m.Invoke(args.Kind, &args)
	}
	// Handle scanner errors if necessary...
}

// NewCommandEventManager creates and returns a new instance of CommandEventManager.
func NewCommandEventManager() *CommandEventManager {
	return &CommandEventManager{
		BaseManager: BaseManager[*Command]{
			lst: make(map[string][]Listener[*Command]),
		},
	}
}

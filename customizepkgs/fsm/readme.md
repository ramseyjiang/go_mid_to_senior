FSM is Finite State Machine.

It has lots of behaviors of FSM in our real lives such as vending machines, elevators, traffic lights, etc.

What is a Finite-State Machine? A finite-state machine (FSM) or simply a state machine, is a mathematical model of
computation. It is an abstract machine that can be in exactly one of a finite number of states at any given time. The
FSM can change from one state to another in response to some inputs; the change from one state to another is called a
transition.

— Wikipedia

An FSM consists of three key elements: an initial state, a list of all possible states, the inputs that trigger
state-transition.

Implement an FSM in an object-oriented way. A class Turnstile is introduced, and it has a propertyState, a method named
ExecuteCmd. The only way to trigger state transition is by calling the method ExecuteCmd.
package mockactivity

import (
	"fmt"
	"net/mail"
)

func apply(tutor, student, what string, all []Application) (Application, error) {
	// behavior: valid activity
	if what != "soccer" {
		return Application{}, fmt.Errorf("%w: %s", ErrInvalidActivity, what)
	}

	// behavior: valid email
	_, err := mail.ParseAddress(tutor)
	if err != nil {
		return Application{}, fmt.Errorf("%w '%s': %s", ErrInvalidTutorEmail, tutor, err.Error())
	}

	// behavior: valid email
	for _, a := range all {
		if a.Student == student {
			return Application{}, ErrAlreadyApplied
		}
	}

	// behavior: cost only for first one
	cost := 10
	for _, a := range all {
		if a.ActivityName == what {
			cost = 0
		}
	}

	// change state
	return Application{
		TutorEmail:   tutor,
		ActivityName: what,
		Student:      student,
		Cost:         cost,
	}, nil
}

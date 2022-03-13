package mockactivity

import "errors"

var (
	ErrInternal          = errors.New("an internal error occurred")
	ErrAlreadyApplied    = errors.New("student already applied")
	ErrInvalidTutorEmail = errors.New("invalid tutor email")
	ErrInvalidActivity   = errors.New("invalid activity")
)

type Application struct {
	TutorEmail   string
	Cost         int
	ActivityName string
	Student      string
}

type ApplicationRepo interface {
	FindAll(tutor string) ([]Application, error)
	Save(application Application) error
}

type EmailSender interface {
	Send(application Application) error
}

type App struct {
	ApplicationRepo ApplicationRepo
	EmailSender     EmailSender
}

func (a *App) Apply(tutor, student, what string) (Application, error) {
	all, err := a.ApplicationRepo.FindAll(tutor) // invoke ApplicationRepo interface
	if err != nil {
		return Application{}, err
	}

	application, err := apply(tutor, student, what, all)
	if err != nil {
		return Application{}, err
	}

	err = a.ApplicationRepo.Save(application) // invoke ApplicationRepo interface
	if err != nil {
		return Application{}, err
	}

	err = a.EmailSender.Send(application) // invoke EmailSender interface
	if err != nil {
		return Application{}, err
	}

	return application, nil
}

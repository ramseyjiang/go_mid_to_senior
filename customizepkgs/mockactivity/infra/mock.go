package infra

import "github.com/ramseyjiang/go_mid_to_senior/customizepkgs/mockactivity"

type EmailSenderMock struct {
	SendFunc func(application mockactivity.Application) error
}

func (e EmailSenderMock) Send(application mockactivity.Application) error {
	return e.SendFunc(application)
}

type ApplicationRepoMock struct {
	FindAllFunc func(tutor string) ([]mockactivity.Application, error)
	SaveFunc    func(application mockactivity.Application) error
}

func (a ApplicationRepoMock) FindAll(tutor string) ([]mockactivity.Application, error) {
	return a.FindAllFunc(tutor)
}

func (a ApplicationRepoMock) Save(application mockactivity.Application) error {
	return a.SaveFunc(application)
}

// emailSenderSuccess := infra.EmailSenderMock{
//		SendFunc: func (application mockactivity.Application) error {
//		return nil
//	},
// }

// emailSenderFail := infra.EmailSenderMock{
//		SendFunc: func (application mockactivity.Application) error {
//		return mockactivity.ErrInternal
//	},
// }

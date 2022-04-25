package mockactivity_test

import (
	"reflect"
	"testing"

	"github.com/ramseyjiang/go_mid_to_senior/customizepkgs/mockactivity"
	"github.com/ramseyjiang/go_mid_to_senior/customizepkgs/mockactivity/infra"
)

func TestApp_Apply(t *testing.T) {
	type deps struct {
		applicationRepo mockactivity.ApplicationRepo
		emailSender     mockactivity.EmailSender
	}

	// mock applicationRepo interface
	applicationRepoEmpty := infra.ApplicationRepoMock{
		FindAllFunc: func(tutor string) ([]mockactivity.Application, error) { return nil, nil },
		SaveFunc:    func(application mockactivity.Application) error { return nil },
	}

	// mock emailSender interface
	emailSenderSuccess := infra.EmailSenderMock{
		SendFunc: func(application mockactivity.Application) error { return nil },
	}

	tests := []struct {
		name    string
		deps    deps
		tutor   string
		student string
		what    string
		want    mockactivity.Application
		wantErr bool
	}{
		{
			name:    "first student",
			deps:    deps{applicationRepo: applicationRepoEmpty, emailSender: emailSenderSuccess},
			tutor:   "some@tutor.com",
			student: "Student Full name",
			what:    "soccer",
			want:    mockactivity.Application{TutorEmail: "some@tutor.com", Cost: 10, ActivityName: "soccer", Student: "Student Full name"},
			wantErr: false,
		},
		{
			name:    "wrong activity",
			deps:    deps{applicationRepo: applicationRepoEmpty, emailSender: emailSenderSuccess},
			tutor:   "some@tutor.com",
			student: "Student Full name",
			what:    "undefined",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &mockactivity.App{ApplicationRepo: tt.deps.applicationRepo, EmailSender: tt.deps.emailSender}
			got, err := a.Apply(tt.tutor, tt.student, tt.what)
			if (err != nil) != tt.wantErr {
				t.Errorf("Apply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Apply() got = %v, want %v", got, tt.want)
			}
		})
	}
}

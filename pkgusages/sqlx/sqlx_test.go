package sqlxpkg

import (
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
)

func TestTrigger(t *testing.T) {
	var tests []struct {
		name string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Trigger()
		})
	}
}

func Test_delUserById(t *testing.T) {
	type args struct {
		db *sqlx.DB
		id int64
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			delUserByID(tt.args.db, tt.args.id)
		})
	}
}

func Test_getAllUsers(t *testing.T) {
	type args struct {
		db *sqlx.DB
	}
	var tests []struct {
		name    string
		args    args
		want    []User
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getAllUsers(tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("getAllUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllUsers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getUserById(t *testing.T) {
	type args struct {
		db *sqlx.DB
		id int64
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getUserByID(tt.args.db, tt.args.id)
		})
	}
}

func Test_insertUser(t *testing.T) {
	type args struct {
		db *sqlx.DB
	}
	var tests []struct {
		name string
		args args
		want int64
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertUser(tt.args.db); got != tt.want {
				t.Errorf("insertUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_updateUserById(t *testing.T) {
	type args struct {
		db *sqlx.DB
		id int64
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateUserByID(tt.args.db, tt.args.id)
		})
	}
}

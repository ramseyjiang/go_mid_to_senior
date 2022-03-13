package main

import (
	"database/sql"
	"testing"
)

func Test_createUsersTable(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createUsersTable(tt.args.db)
		})
	}
}

func Test_delAllRows(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			delAllRows(tt.args.db)
		})
	}
}

func Test_delRowByID(t *testing.T) {
	type args struct {
		db     *sql.DB
		userID int64
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			delRowByID(tt.args.db, tt.args.userID)
		})
	}
}

func Test_getAllRows(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getAllRows(tt.args.db)
		})
	}
}

func Test_getRowByID(t *testing.T) {
	type args struct {
		db     *sql.DB
		userID int64
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getRowByID(tt.args.db, tt.args.userID)
		})
	}
}

func Test_insertRow(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	var tests []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := insertRow(tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("insertRow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("insertRow() got = %v, want %v", got, tt.want)
			}
		})
	}
}

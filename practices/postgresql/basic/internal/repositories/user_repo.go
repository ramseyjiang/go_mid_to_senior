package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"pgtest.com/m/v2/internal/models"
)

type UserRepository interface {
	CreateTable() error
	CreateUser(user *models.User) (int, error)
	GetUserByID(id int) (*models.User, error)
}

type PostgresUserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) CreateTable() error {
	user := &models.User{}
	return user.CreateTable(r.db)
}

func (r *PostgresUserRepository) CreateUser(user *models.User) (int, error) {
	return user.CreateUser(r.db)
}

func (r *PostgresUserRepository) GetUserByID(id int) (*models.User, error) {
	type result struct {
		user *models.User
		err  error
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resChan := make(chan result, 1)

	go func() {
		user := &models.User{}
		userInfo, err := user.GetUserByID(r.db, id)
		resChan <- result{user: userInfo, err: err}
	}()

	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("query timeout after 3 seconds")
	case res := <-resChan:
		return res.user, res.err
	}
}

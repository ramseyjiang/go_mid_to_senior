package migrations

import (
	"golang_learn/projects/users/config"
	"golang_learn/projects/users/entity"
)

func getModels() []interface{} {
	return []interface{}{&entity.User{}}
}

func MigrateTable() {
	db := config.GetDB()
	_ = db.AutoMigrate(getModels()...)
}

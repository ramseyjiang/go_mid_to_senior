package migrations

import (
	"github.com/ramseyjiang/go_mid_to_senior/projects/users/config"
	"github.com/ramseyjiang/go_mid_to_senior/projects/users/entity"
)

func getModels() []interface{} {
	return []interface{}{&entity.User{}}
}

func MigrateTable() {
	db := config.GetDB()
	_ = db.AutoMigrate(getModels()...)
}

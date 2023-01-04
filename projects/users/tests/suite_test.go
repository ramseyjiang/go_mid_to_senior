package tests

import (
	"database/sql"
	"os"
	"testing"

	"github.com/ramseyjiang/go_mid_to_senior/projects/users/config"
	"github.com/ramseyjiang/go_mid_to_senior/projects/users/models"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type SuiteTest struct {
	suite.Suite
	db *gorm.DB
}

func TestSuite(t *testing.T) {
	_ = os.Setenv("DB_HOST", "127.0.0.1")
	_ = os.Setenv("DB_PORT", "3306")
	_ = os.Setenv("DB_USER", "root")
	_ = os.Setenv("DB_PASSWORD", "12345678")
	_ = os.Setenv("DB_DATABASE", "go_web")
	defer func() {
		err := os.Unsetenv("DB_HOST")
		if err != nil {
			return
		}
	}()
	defer func() {
		err := os.Unsetenv("DB_PORT")
		if err != nil {
			return
		}
	}()
	defer func() {
		err := os.Unsetenv("DB_USER")
		if err != nil {
			return
		}
	}()
	defer func() {
		err := os.Unsetenv("DB_PASS")
		if err != nil {
			return
		}
	}()
	defer func() {
		err := os.Unsetenv("DB_DATABASE")
		if err != nil {
			return
		}
	}()

	suite.Run(t, new(SuiteTest))
}

func getModels() []interface{} {
	return []interface{}{&models.User{}}
}

// Setup data value
func (t *SuiteTest) SetupSuite() {
	config.ConnectGorm()
	t.db = config.GetDB()

	// Migrate Table
	for _, val := range getModels() {
		_ = t.db.AutoMigrate(val)
	}
}

// Run After All Test Done
func (t *SuiteTest) TearDownSuite() {
	sqlDB, _ := t.db.DB()
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {
			return
		}
	}(sqlDB)

	// Drop Table, if open the comment, it will drop users table during tests run.
	// for _, val := range getModels() {
	// 	_ = t.data.Migrator().DropTable(val)
	// }
}

// Run Before a Test
func (t *SuiteTest) SetupTest() {
}

// Run After a Test
func (t *SuiteTest) TearDownTest() {
}

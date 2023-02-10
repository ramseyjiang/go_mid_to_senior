package dbconn

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestConfig(t *testing.T) {
	confTest := &DBConfig{
		Username: "user1",
		Password: "pass1",
		Host:     "localhost",
		Port:     5432,
		DBName:   "db1",
	}
	ConfManager.SetConfig("config test", confTest)

	configDev := &DBConfig{
		Username: "user2",
		Password: "pass2",
		Host:     "localhost",
		Port:     5433,
		DBName:   "db2",
	}
	ConfManager.SetConfig("config dev", configDev)

	config, err := ConfManager.GetConfig("config test")
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(t, confTest.Username, config.Username)
	assert.Equal(t, confTest.Port, config.Port)

	config, err = ConfManager.GetConfig("config dev")
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(t, configDev.Username, config.Username)
	assert.Equal(t, configDev.Port, config.Port)
}

package dbconn

import (
	"fmt"
	"sync"
)

// DBConfig represents a database configuration
type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	DBName   string
}

// ConfigManager represents a database configuration manager
type ConfigManager struct {
	configs map[string]*DBConfig
	once    sync.Once
}

// ConfManager is a singleton instance of ConfigManager
var ConfManager ConfigManager

// GetConfig returns a database configuration for a given key
func (c *ConfigManager) GetConfig(key string) (*DBConfig, error) {
	c.once.Do(func() {
		c.configs = make(map[string]*DBConfig)
	})
	config, ok := c.configs[key]
	if !ok {
		return nil, fmt.Errorf("config for key '%s' not found", key)
	}
	return config, nil
}

// SetConfig sets a database configuration for a given key
func (c *ConfigManager) SetConfig(key string, config *DBConfig) {
	c.once.Do(func() {
		c.configs = make(map[string]*DBConfig)
	})
	c.configs[key] = config
}

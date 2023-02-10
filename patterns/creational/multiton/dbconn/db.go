package dbconn

import (
	"fmt"
	"sync"
)

// In this example, we have a database configuration manager represented by the ConfigManager struct. The manager holds a map of database configurations, with each configuration identified by a unique key. The sync.Once type is used to ensure that the map of configurations is only created once, even if multiple goroutines attempt to access the map concurrently. The Instance variable is a singleton instance of ConfigManager, and the GetConfig and SetConfig methods allow for retrieving and setting database configurations in the map, respectively.

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

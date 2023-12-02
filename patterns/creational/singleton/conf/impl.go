package conf

import "sync"

type Config struct {
	settings map[string]string
}

var instance *Config
var once sync.Once

func GetInstance() *Config {
	once.Do(func() {
		instance = &Config{
			settings: make(map[string]string),
		}
	})
	return instance
}

func (c *Config) Set(key, value string) {
	c.settings[key] = value
}

func (c *Config) Get(key string) string {
	return c.settings[key]
}

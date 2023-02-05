package conf

import "sync"

type config struct {
	settings map[string]string
}

var instance *config
var once sync.Once

func GetInstance() *config {
	once.Do(func() {
		instance = &config{
			settings: make(map[string]string),
		}
	})
	return instance
}

func (c *config) Set(key, value string) {
	c.settings[key] = value
}

func (c *config) Get(key string) string {
	return c.settings[key]
}

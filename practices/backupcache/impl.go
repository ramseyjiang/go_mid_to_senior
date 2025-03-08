package backupcache

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type CacheData struct {
	Data string `json:"data"`
}

func GetCacheData(key string) (*CacheData, error) {
	// attempt to load the backup cache from the local file
	filePath := "/tmp/backup_cache/" + key + ".json"
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var cacheData CacheData
	err = json.Unmarshal(data, &cacheData)
	if err != nil {
		return nil, err
	}

	return &cacheData, nil
}

// SaveCacheData is used to store data in the "cd /tmp/backup_cache folder"
func SaveCacheData(key string, data *CacheData) error {
	// Store data in the local "cd /tmp/backup_cache folder"
	filePath := "/tmp/backup_cache/" + key + ".json"
	err := os.MkdirAll("/tmp/backup_cache/", 0755) // Create folder if it doesn't exist
	if err != nil {
		return err
	}
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}
	return nil
}

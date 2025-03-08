package backupcache

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCacheData(t *testing.T) {
	key := "test_key"
	data := &CacheData{Data: "test_data"}
	filePath := "/tmp/backup_cache/" + key + ".json"

	// 清除测试文件
	os.Remove(filePath)

	// 测试保存数据
	err := SaveCacheData(key, data)
	assert.NoError(t, err)

	// 验证文件存在
	_, err = os.Stat(filePath)
	assert.NoError(t, err)

	// 测试读取数据
	cachedData, err := GetCacheData(key)
	assert.NoError(t, err)
	assert.Equal(t, data, cachedData)
}

func TestCacheDataWithTable(t *testing.T) {
	tests := []struct {
		name       string
		key        string
		data       *CacheData
		setup      func(key string, data *CacheData) error
		wantErrGet bool
		wantData   *CacheData
	}{
		{
			name: "Test successful",
			key:  "normal_key",
			data: &CacheData{Data: "normal_data"},
			setup: func(key string, data *CacheData) error {
				return SaveCacheData(key, data)
			},
			wantErrGet: false,
			wantData:   &CacheData{Data: "normal_data"},
		},
		{
			name: "File does not exist",
			key:  "missing_key",
			data: nil,
			setup: func(key string, data *CacheData) error {
				// 不执行任何操作，文件不存在
				return nil
			},
			wantErrGet: true,
			wantData:   nil,
		},
		{
			name: "Invalid JSON",
			key:  "invalid_json_key",
			data: nil,
			setup: func(key string, data *CacheData) error {
				filePath := "/tmp/backup_cache/" + key + ".json"
				return os.WriteFile(filePath, []byte("invalid json"), 0644)
			},
			wantErrGet: true,
			wantData:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filePath := "/tmp/backup_cache/" + tt.key + ".json"
			// 清除测试文件
			os.Remove(filePath)

			if tt.setup != nil {
				err := tt.setup(tt.key, tt.data)
				assert.NoError(t, err)
			}

			got, err := GetCacheData(tt.key)
			if tt.wantErrGet {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantData, got)
			}
		})
	}
}

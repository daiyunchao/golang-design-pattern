package singleton

import "sync"

type ConfigManager struct {
	ConfigCount *int
}

var configManager *ConfigManager
var once sync.Once

func GetConfigManager() *ConfigManager {
	once.Do(func() {
		count := 100
		configManager = &ConfigManager{
			ConfigCount: &(count),
		}
	})
	return configManager
}

func (cm *ConfigManager) GetConfigCount() *int {
	return cm.ConfigCount
}

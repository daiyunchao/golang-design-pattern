package singleton

import (
	"testing"
)

func TestGetConfigManager(t *testing.T) {
	configManager := GetConfigManager()
	count1 := configManager.GetConfigCount()
	count2 := configManager.GetConfigCount()
	if count1 != count2 {
		t.Errorf("GetConfigManager failed, expected %d got %d", count1, count2)
	}
	t.Logf("GetConfigManager success, expected %d got %d", count1, count2)
}

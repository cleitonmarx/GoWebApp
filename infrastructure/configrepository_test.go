package infrastructure

import (
	"reflect"
	"testing"

	"github.com/cleitonmarx/gowebapp/config"
)

func TestGetSystemConfiguration_WithValidFilePath_NotEmptyConfig(t *testing.T) {
	configRepository := NewConfigFileRepository("../gowebapp.json")
	configuration, err := configRepository.GetSystemConfiguration()
	if err != nil {
		t.Error(err)
	}

	currentConfig, err := configuration.GetCurrentEnvironmentConfig()
	if err != nil {
		t.Error(err)
	}

	emptyConfig := config.Configuration{}
	if reflect.DeepEqual(currentConfig, emptyConfig) {
		t.Error("currentConfig should not be empty")
	}

}

func TestGetSystemConfiguration_WithInvalidFilePath_ReturnAnError(t *testing.T) {
	configRepository := NewConfigFileRepository("InvalidFile.config")
	currentConfig, err := configRepository.GetSystemConfiguration()
	if err == nil {
		t.Error("Error should not be nil")
	}
	emptyConfig := config.Configuration{}
	if !reflect.DeepEqual(currentConfig, emptyConfig) {
		t.Error("currentConfig should be empty")
	}

}

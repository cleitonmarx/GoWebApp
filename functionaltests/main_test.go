package functionaltests

import (
	"net/http"
	"testing"

	"github.com/cleitonmarx/gowebapp/config"
	"github.com/cleitonmarx/gowebapp/infrastructure"
)

func Test_RootPath_Get200Response(t *testing.T) {
	if testing.Short() {
		t.Skip("Functional tests, run just for integration")
	}

	envConfig, err := getCurrentConfig()
	if err != nil {
		t.Error(err.Error())
	}

	client := http.DefaultClient
	response, err := client.Get("http://" + envConfig.HTTPServer.GetFormatedAddress() + "/")
	if err != nil {
		t.Error(err.Error())
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Status 200 expected, %d actual", response.StatusCode)
	}
}
func Test_VersionPath_Get200Response(t *testing.T) {
	if testing.Short() {
		t.Skip("Functional tests, run just for integration")
	}

	envConfig, err := getCurrentConfig()
	if err != nil {
		t.Error(err.Error())
	}

	client := http.DefaultClient
	response, err := client.Get("http://" + envConfig.HTTPServer.GetFormatedAddress() + "/_version")
	if err != nil {
		t.Error(err.Error())
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Status 200 expected, %d actual", response.StatusCode)
	}
}

func getCurrentConfig() (config.EnvironmentConfig, error) {
	result := config.EnvironmentConfig{}
	confRepo := infrastructure.NewConfigFileRepository("../gowebapp.json")
	systemConfig, err := confRepo.GetSystemConfiguration()
	if err != nil {
		return result, err
	}

	currentConfig, err := systemConfig.GetCurrentEnvironmentConfig()
	if err != nil {
		return result, err
	}

	return currentConfig, nil
}

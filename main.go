package main

import (
	"os"
	"strings"
	"time"

	"github.com/cleitonmarx/gowebapp/config"
	"github.com/cleitonmarx/gowebapp/infrastructure"
	"github.com/cleitonmarx/gowebapp/server"
)

func main() {
	var configFileRepository config.Repository
	configFileRepository = infrastructure.NewConfigFileRepository(
		strings.Join([]string{os.Getenv("GOPATH"), "/src/github.com/cleitonmarx/gowebapp/gowebapp.json"}, ""),
	)
	currentConfig := getCurrentConfig(configFileRepository)
	appServer := server.New(currentConfig)
	appServer.Init()
	appServer.Run()
}

//getCurrentConfig gets the configuration variables from the current environment
func getCurrentConfig(configRepository config.Repository) config.EnvironmentConfig {
	systemConfig, err := configRepository.GetSystemConfiguration()
	handleError(err)

	// ---- Pause app for network to initialize during production ----
	if systemConfig.CurrentEnvironment == "PRODUCTION" {
		time.Sleep(time.Second * 1)
	}

	currentConfig, err := systemConfig.GetCurrentEnvironmentConfig()
	handleError(err)

	return currentConfig
}

//handleError stops the program execution
func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

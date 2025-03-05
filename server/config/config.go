package config

import (
	"alter-io-go/helpers/logger"
	"log/slog"
	"sync"

	"github.com/spf13/viper"
)

var (
	appConfig *AppConfig
	lock      = &sync.Mutex{}
)

type AppConfig struct {
	App struct {
		Port       uint16   `toml:"port"`
		CorsOrigin []string `toml:"corsorigin"`
	} `toml:"app"`
	Database struct {
		Name              string `toml:"name"`
		Username          string `toml:"username"`
		Password          string `toml:"password"`
		Port              uint16 `toml:"port"`
		Address           string `toml:"address"`
		Driver            string `toml:"driver"`
		MaxConns          int32  `toml:"maxconns"`
		MinConns          int32  `toml:"minconns"`
		MaxConnLifeTime   int32  `toml:"maxconnlifetime"`
		MaxConnIdleTime   int32  `toml:"maxconnidletime"`
		HealthCheckPeriod int32  `toml:"healthcheckperiod"`
		ConnTimeout       int32  `toml:"conntimeout"`
	} `toml:"database"`
	JWT struct {
		Secret  string `toml:"secret"`
		Issuer  string `toml:"issuer"`
		Subject string `toml:"subject"`
	} `toml:"jwt"`
}

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		var err error
		appConfig, err = loadConfig()
		if err != nil {
			logger.Get().With(slog.String("error", err.Error())).Error("Failed to load config")
		}
	}

	return appConfig
}
func loadConfig() (*AppConfig, error) {
	viper.AddConfigPath("./config/")
	viper.SetConfigType("toml")
	viper.SetConfigName("app")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Get().With(slog.String("error", err.Error())).Error("Config file not found")
			return nil, err
		}
	}

	var config AppConfig
	err := viper.Unmarshal(&config)
	if err != nil {
		logger.Get().With(slog.String("error", err.Error())).Error("Failed to unmarshal config")
		return nil, err
	}

	// Debugging the loaded config
	logger.Get().With().Info("Configuration loaded")

	return &config, nil
}

package configs

import (
	"time"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Handler  *HandlerConfig
		DataBase *DatabaseConfig
	}

	DatabaseConfig struct {
		Host     string
		Port     string
		Username string
		Password string
		DBName   string
		SSLMode  string
	}

	HandlerConfig struct {
		UrlId          string
		RequestTimeout time.Duration
	}
)

func InitConfig() (*Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Config{
		DataBase: &DatabaseConfig{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Username: viper.GetString("db.username"),
			DBName:   viper.GetString("db.dbname"),
			SSLMode:  viper.GetString("db.sslmode"),
		},
		Handler: &HandlerConfig{
			RequestTimeout: viper.GetDuration("handler.requestTimeout"),
			UrlId:          viper.GetString("handler.urlId"),
		},
	}, nil
}

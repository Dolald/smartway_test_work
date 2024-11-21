package configs

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type (
	Config struct {
		Handler  *HandlerConfig
		DataBase *DatabaseConfig
		Server   *ServerConfig
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

	ServerConfig struct {
		Port           string
		MaxHeaderBytes int
		ReadTimeout    time.Duration
		WriteTimeout   time.Duration
	}
)

func InitConfig() (*Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error in InitConfig: ReadInConfig: %w", err)
	}

	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error in InitConfig: Load failed: %w", err)
	}

	return &Config{
		DataBase: &DatabaseConfig{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Username: viper.GetString("db.username"),
			DBName:   viper.GetString("db.dbname"),
			SSLMode:  viper.GetString("db.sslmode"),
			Password: os.Getenv("DB_PASSWORD"),
		},
		Handler: &HandlerConfig{
			RequestTimeout: viper.GetDuration("handler.requestTimeout"),
			UrlId:          viper.GetString("handler.urlId"),
		},
		Server: &ServerConfig{
			Port:           viper.GetString("server.port"),
			MaxHeaderBytes: viper.GetInt("server.maxHeaderBytes"),
			ReadTimeout:    viper.GetDuration("server.readTimeout"),
			WriteTimeout:   viper.GetDuration("server.writeTimeout"),
		},
	}, nil
}

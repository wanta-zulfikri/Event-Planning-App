package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Database           Database `mapstructure:"Database"`
	JWTConfig          JWTConfig
	GoogleCloudStorage GoogleCloudStorage `mapstructure:"GoogleCloudStorage"`
	Server             Server             `mapstructure:"Server"`
}

type Database struct {
	Username string `mapstructure:"Username"`
	Password string `mapstructure:"Password"`
	Host     string `mapstructure:"Host"`
	Port     string `mapstructure:"Port"`
	Name     string `mapstructure:"Name"`
}

type JWTConfig struct {
	Secret        string
	SigningMethod string
}

type Server struct {
	Port string `mapstructure:"Port"`
}

type GoogleCloudStorage struct {
	Credential string `mapstructure:"Credential"`
	ProjectID  string `mapstructure:"ProjectID"`
	Bucketname string `mapstructure:"Bucketname"`
	Path       string `mapstructure:"Path"`
}

func InitConfiguration() (*Config, error) {
	LoadEnv()

	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Println("Viper failed to read config, attempting to read from environment variables...")
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	if config.Database.Host == "" {
		var defaultConfig Config
		defaultConfig.Database.Host = os.Getenv("Host")
		defaultConfig.Database.Name = os.Getenv("Name")
		defaultConfig.Database.Port = os.Getenv("Port")
		defaultConfig.Database.Username = os.Getenv("Username")
		defaultConfig.Database.Password = os.Getenv("Password")
		return &defaultConfig, nil
	}

	return &config, nil
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, will use environment variables instead.")
	}
}

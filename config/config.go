package config

import "github.com/spf13/viper"

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
	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	viper.AutomaticEnv()
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

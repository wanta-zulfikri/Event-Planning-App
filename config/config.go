package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/wanta-zulfikri/Event-Planning-App/config/common"
)

type Configuration struct {
	Port     string
	Database struct {
		Driver   string
		Host     string
		Name     string
		Address  string
		Port     string
		Username string
		Password string
	}
}

var lock = &sync.Mutex{}
var appConfig *Configuration

func GetConfiguration() *Configuration {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = InitConfiguration()
	}

	return appConfig
}

func InitConfiguration() *Configuration {

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	var defaultConfig Configuration
	defaultConfig.Port = os.Getenv("AppPort")
	defaultConfig.Database.Host = os.Getenv("Host")
	defaultConfig.Database.Port = os.Getenv("Port")
	defaultConfig.Database.Username = os.Getenv("Username")
	defaultConfig.Database.Password = os.Getenv("Password")
	defaultConfig.Database.Name = os.Getenv("Name")
	common.JWTSecret = os.Getenv("JWTSecret")
	common.Credential = os.Getenv("Credential")
	common.ProjectID = os.Getenv("ProjectID")
	common.BucketName = os.Getenv("BucketName")
	common.Path = os.Getenv("Path")
	common.MidstransServerKey = os.Getenv("MidstransServerKey")

	return &defaultConfig

}

type SenderConfig struct {
	Email     string
	Password  string
	Phone     string
	Name      string
	Address   string
	Slogan    string
	Twitter   string
	Instagram string
	Facebook  string
}

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

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// env, err := strconv.Atoi(os.Getenv("Environment"))
	// if err != nil {
	// 	log.Fatal("Error parsing Unit: ", err)
	// }

	// expiryDuration, err := strconv.Atoi(os.Getenv("ExpiryDuration"))
	// if err != nil {
	// 	log.Fatal("Error parsing ExpiryDuration: ", err)
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
	// common.MIDTRANS_CLIENT_KEY = os.Getenv("MIDTRANS_CLIENT_KEY")
	// common.MIDTRANS_SERVER_KEY = os.Getenv("MIDTRANS_SERVER_KEY")
	// common.Environment = env
	// common.URLHandler = os.Getenv("URLHandler")
	// common.ExpiryDuration = expiryDuration
	// common.Unit = os.Getenv("Unit")

	return &defaultConfig

}

type NSQConfig struct {
	Host     string `mapstructure:"HOST"`
	Port     string `mapstructure:"PORT"`
	Topic    string `mapstructure:"TOPIC"`
	Channel  string `mapstructure:"CHANNEL"`
	Topic2   string `mapstructure:"TOPIC2"`
	Channel2 string `mapstructure:"CHANNEL2"`
	Topic3   string `mapstructure:"TOPIC3"`
	Channel3 string `mapstructure:"CHANNEL3"`
	Topic4   string `mapstructure:"TOPIC4"`
	Channel4 string `mapstructure:"CHANNEL4"`
} 


type SenderConfig struct {
	Email     string `mapstructure:"EMAIL"`
	Password  string `mapstructure:"PASSWORD"`
	Phone     string `mapstructure:"PHONE"`
	Name      string `mapstructure:"NAME"`
	Address   string `mapstructure:"ADDRESS"`
	Slogan    string `mapstructure:"SLOGAN"`
	Twitter   string `mapstructure:"TWTR"`
	Instagram string `mapstructure:"IG"`
	Facebook  string `mapstructure:"FB"`
} 


type (
	Data struct {
		Invoice       string `json:"invoice"`
		Total         int    `json:"total"`
		Name          string `json:"name"`
		Email         string `json:"email"`
		PaymentCode   string `json:"payment_code"`
		PaymentMethod string `json:"payment_method"`
		Expire        string `json:"expire"`
	}
)
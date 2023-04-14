package config

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

var (
	JWTKEY                 string
	CloudinaryName         string
	CloudinaryApiKey       string
	CloudinaryApiScret     string
	CloudinaryUploadFolder string
)

type AppConfig struct {
	DB_USER string
	DB_PASS string
	DB_HOST string
	DB_PORT int
	DB_NAME string
}

func InitConfig() *AppConfig {
	return ReadEnv()
}

func ReadEnv() *AppConfig {
	app := AppConfig{}
	isRead := true

	if val, found := os.LookupEnv("DB_USER"); found {
		app.DB_USER = val
		isRead = false
	}
	if val, found := os.LookupEnv("DB_PASS"); found {
		app.DB_PASS = val
		isRead = false
	}
	if val, found := os.LookupEnv("DB_HOST"); found {
		app.DB_HOST = val
		isRead = false
	}
	if val, found := os.LookupEnv("DB_PORT"); found {
		cnv, _ := strconv.Atoi(val)
		app.DB_PORT = cnv
		isRead = false
	}
	if val, found := os.LookupEnv("DB_NAME"); found {
		app.DB_NAME = val
		isRead = false
	}

	if val, found := os.LookupEnv("JWT_KEY"); found {
		JWTKEY = val
		isRead = false
	}

	if val, found := os.LookupEnv("CLOUDINARY_CLOUD_NAME"); found {
		CloudinaryName = val
		isRead = false
	}
	if val, found := os.LookupEnv("CLOUDINARY_API_KEY"); found {
		CloudinaryApiKey = val
		isRead = false
	}
	if val, found := os.LookupEnv("CLOUDINARY_API_SECRET"); found {
		CloudinaryApiScret = val
		isRead = false
	}
	if val, found := os.LookupEnv("CLOUDINARY_UPLOAD_FOLDER"); found {
		CloudinaryUploadFolder = val
		isRead = false
	}

	if isRead {
		viper.AddConfigPath(".")
		viper.SetConfigName("local")
		viper.SetConfigType("env")

		err := viper.ReadInConfig()
		if err != nil {
			log.Println("error read config : ", err.Error())
			return nil
		}

		app.DB_USER = viper.Get("DB_USER").(string)
		app.DB_PASS = viper.Get("DB_PASS").(string)
		app.DB_HOST = viper.Get("DB_HOST").(string)
		app.DB_PORT, _ = strconv.Atoi(viper.Get("DB_PORT").(string))
		app.DB_NAME = viper.Get("DB_NAME").(string)

		JWTKEY = viper.Get("JWT_KEY").(string)

		CloudinaryName = viper.Get("CLOUDINARY_CLOUD_NAME").(string)
		CloudinaryApiKey = viper.Get("CLOUDINARY_API_KEY").(string)
		CloudinaryApiScret = viper.Get("CLOUDINARY_API_SECRET").(string)
		CloudinaryUploadFolder = viper.Get("CLOUDINARY_UPLOAD_FOLDER").(string)

	}
	return &app
}
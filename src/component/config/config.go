package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

const (
	DefaultFileExtension = "yaml"
	DefaultFileName      = "main"
	DefaultFilePath      = "./configs"
)

type Config struct {
	configurator *viper.Viper

	DB     *Database
	Limits *Limits

	filePath      string
	fileName      string
	fileExtension string
}

type Limits struct {
	Login    int
	Password int
	IP       int
}

type Database struct {
	Host     string
	User     string
	Password string
	Dbname   string
	Port     string
	Charset  string
	Sslmode  string
}

func NewConfig(configPath string, fileName string, fileExtension string) *Config {
	var configData = Config{
		filePath:      configPath,
		fileName:      fileName,
		fileExtension: fileExtension,
	}

	if configData.fileExtension == "" {
		configData.fileExtension = DefaultFileExtension
	}

	if configData.fileName == "" {
		configData.fileName = DefaultFileName
	}

	if configData.filePath == "" {
		configData.filePath = DefaultFilePath
	}

	configData.addConfigurator()

	return &configData
}

func (config *Config) Configurator() *viper.Viper {
	return config.configurator
}

func (config *Config) Fill() {
	config.configurator.AddConfigPath(config.filePath)

	if err := config.configurator.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}

	//merge local config
	var localFullNameLocalFile = fmt.Sprintf("%s/%s.%s", config.filePath, config.fileNameLocal(), config.fileExtension)
	if config.isExistsFile(localFullNameLocalFile) {
		config.configurator.SetConfigName(config.fileNameLocal())
		if err := config.configurator.MergeInConfig(); err != nil {
			log.Fatal(err)
		}
	}

	if err := config.unSerialize(); err != nil {
		log.Fatal(err)
	}
}

func (config *Config) isExistsFile(fullFileName string) bool {
	if _, err := os.Stat(fullFileName); os.IsNotExist(err) {
		return false
	}

	return true
}

func (config *Config) addConfigurator() {
	config.configurator = viper.New()
	config.configurator.SetConfigType(config.fileExtension)
	config.configurator.SetConfigName(config.fileName)
}

func (config *Config) unSerialize() error {
	err := config.configurator.Unmarshal(&config)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func (config Config) fileNameLocal() string {
	return config.fileName
}

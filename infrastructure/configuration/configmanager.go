package configuration

import (
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
)

type configurationManager interface {
	GetServerConfig() ServerConfig
	GetDatabaseConfig() DatabaseConfig
}

type configurationManagerImp struct {
	applicationConfig ApplicationConfig
}

func (configurationManager *configurationManagerImp) GetServerConfig() ServerConfig {
	return configurationManager.applicationConfig.Server
}

func (configurationManager *configurationManagerImp) GetDatabaseConfig() DatabaseConfig {
	return configurationManager.applicationConfig.Database
}



var instance *configurationManagerImp

const (
	configPath = "./infrastructure/resource"
	serverPort = ":6161"
	configType = "yaml"
	configName = "application"
)

func NewConfigurationManager() configurationManager {
	if instance != nil {
		return instance
	}

	env := os.Getenv("ACTIVE_PROFILE")
	if env == "" {
		log.Print("**** ACTIVE_PROFILE is empty, default it will be used as 'dev' ****")
		env = "dev"
	}

	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)
	viper.SetTypeByDefaultValue(true)
	viper.SetDefault("server.port", serverPort)
	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		readConf(env)
		logrus.WithField("file", e.Name).Warn("Config file changed")
	})

	instance := readConf(env)

	return &configurationManagerImp{
		applicationConfig: instance,
	}
}

func readConf(env string) ApplicationConfig {
	readConfigErr := viper.ReadInConfig()
	if readConfigErr != nil {
		log.Panicf("Couldn't load application configuration, cannot start. Error details: %s", readConfigErr.Error())
	}
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)
	viper.SetTypeByDefaultValue(true)
	mergeConfigErr := viper.MergeInConfig()
	if mergeConfigErr != nil {
		log.Panicf("Couldn't load application configuration, cannot start. Error details: %s", mergeConfigErr.Error())
	}
	var conf ApplicationConfig
	c := viper.Sub(env)
	unMarshalErr := c.Unmarshal(&conf)
	unMarshalSubErr := c.Unmarshal(&conf)

	if unMarshalErr != nil {
		log.Panicf("Configuration cannot deserialize. Terminating. Error details: %s", unMarshalErr.Error())
	}
	if unMarshalSubErr != nil {
		log.Panicf("Configuration cannot deserialize. Terminating. Error details: %s", unMarshalSubErr.Error())
	}

	logrus.WithField("configuration", conf).Debug("Configuration changed")

	return conf
}

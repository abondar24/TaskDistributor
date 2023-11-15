package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server   ServerConfig
	Broker   BrokerConfig
	Database DatabaseConfig
	RPC      RpcServerConfig
}

type ServerConfig struct {
	Port int
}

type BrokerConfig struct {
	Username  string
	Password  string
	Host      string
	Port      int
	QueueName string
}

type DatabaseConfig struct {
	Username      string
	Password      string
	Host          string
	Port          int
	Driver        string
	Database      string
	MigrationPath string
}

type RpcServerConfig struct {
	Host string
	Port int
}

func ReadConfig() *Config {
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s\n", err)
	}

	var conf Config
	err = viper.Unmarshal(&conf)
	if err != nil {
		log.Fatalf("Unable to decode into struct: %s\n", err)
	}

	return &conf
}

package config

type Config struct {
	Server   ServerConfig
	Broker   BrokerConfig
	Database DatabaseConfig
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
	Username string
	Password string
	Host     string
	Port     int
}

package models

type ServerConfig struct {
	Host string
	Port string
}

type DatabaseConfig struct {
	Host string  
	User string 
	Password string
	DBName string
	Port string
	SSLMode string
}

type Config struct {
	Server ServerConfig
	Database DatabaseConfig
}

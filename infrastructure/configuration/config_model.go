package configuration

type ApplicationConfig struct {
	Server    ServerConfig    `yaml:"server"`
	Database  DatabaseConfig  `yaml:"database"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type DatabaseConfig struct {
	UserName          string
	Password          string
	Database          string
	Host              string
	MaxOpenConnection int
	MaxIdleConnection int
	Port              int
}
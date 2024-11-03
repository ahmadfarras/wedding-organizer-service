package model

type Config struct {
	Service     string         `mapstructure:"SERVICE"`
	Version     string         `mapstructure:"VERSION"`
	Environment string         `mapstructure:"ENVIRONMENT"`
	Port        string         `mapstructure:"PORT"`
	Database    DatabaseConfig `mapstructure:"DATABASE"`
}

type DatabaseConfig struct {
	MySQL MySQLConfig `mapstructure:"MYSQL"`
}

type MySQLConfig struct {
	Host                  string `mapstructure:"HOST"`
	Port                  int    `mapstructure:"PORT"`
	User                  string `mapstructure:"USER"`
	Password              string `mapstructure:"PASSWORD"`
	Name                  string `mapstructure:"NAME"`
	MaxConnections        int    `mapstructure:"MAX_CONNECTIONS"`
	MaxIdleConnections    int    `mapstructure:"MAX_IDLE_CONNECTIONS"`
	ConnectionTimeout     int    `mapstructure:"CONNECTION_TIMEOUT"`
	ReadTimeout           int    `mapstructure:"READ_TIMEOUT"`
	WriteTimeout          int    `mapstructure:"WRITE_TIMEOUT"`
	IdleTimeout           int    `mapstructure:"IDLE_TIMEOUT"`
	MaxLifetimeConnection int    `mapstructure:"MAX_LIFETIME_CONNECTION"`
	MigrationsPath        string `mapstructure:"MIGRATIONS_PATH"`
}

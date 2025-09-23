package config

type DatabaseConfig struct {
	Username, Password, Host, Port, Database string
	MaxAttempts, SecondsToConnect            int
}

package config

type DatabaseConfig struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DatabaseName string
}

func MakeDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Host:         GetEnvOrDefault("DB_HOST", "0.0.0.0"),
		Port:         GetEnvOrDefault("DB_HOST", "3306"),
		Username:     GetEnvOrDefault("DB_HOST", "root"),
		Password:     GetEnvOrDefault("DB_HOST", "password"),
		DatabaseName: GetEnvOrDefault("DB_HOST", "public"),
	}
}

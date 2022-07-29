package config

type HTTPConfig struct {
	Host string
	Port string
}

func MakeHttpConfig() HTTPConfig {
	return HTTPConfig{
		Host: GetEnvOrDefault("HT_HOST", "0.0.0.0"),
		Port: GetEnvOrDefault("HT_PORT", "3000"),
	}
}

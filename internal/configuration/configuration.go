package configuration

type Config struct {
	URL string
}

func GetConfig() (*Config, error) {
	config := &Config{
		URL: "localhost:8080",
	}

	return config, nil
}

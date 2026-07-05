package config

type Config struct {
	Server string
}

func GetConfig() (*Config, error) {

	var config Config
	config = Config{
		Server: ":8888",
	}

	return &config, nil
}

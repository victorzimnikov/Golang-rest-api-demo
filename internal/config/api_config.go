package config

type ApiConfig struct {
	ServerPort  string
	DatabaseURL string
}

func LoadApiConfig() (ApiConfig, error) {
	serverPort, err := requiredEnv("SERVER_PORT")
	if err != nil {
		return ApiConfig{}, err
	}

	databaseURL, err := requiredEnv("DATABASE_URL")
	if err != nil {
		return ApiConfig{}, err
	}

	return ApiConfig{
		ServerPort:  serverPort,
		DatabaseURL: databaseURL,
	}, nil
}

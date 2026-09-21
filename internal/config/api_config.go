package config

type ApiConfig struct {
	ServerPort string
}

func LoadApiConfig() (ApiConfig, error) {
	serverPort, err := requiredEnv("SERVER_PORT")
	if err != nil {
		return ApiConfig{}, err
	}

	return ApiConfig{
		ServerPort: serverPort,
	}, nil
}

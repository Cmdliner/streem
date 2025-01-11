package config

import (
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int
		URI string
	}
	MongoDB struct {
		URI string
		Name string
	}
	JWT struct {
		Secret string
		ExpirationHours int
	}
}

func Load() (*Config, error) {
	isProdEnvironment := os.Getenv("GO_ENV") == "PRODUCTION"

	if isProdEnvironment {
		var cfg Config

		envPort, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			return nil, err
		}
		envServerUri := os.Getenv("SERVER_URI")
		envMongoURI := os.Getenv("MONGO_URI")
		envDbName := os.Getenv("DB_NAME")
		envJwtSecret := os.Getenv("JWT_SECRET")
		jwtExpiry, err := strconv.Atoi(os.Getenv("JWT_EXPIRES_HRS"))
		if err != nil {
			return nil, err
		}

		// Set the  into config keys if they pass the checks above
		cfg.Server.URI = envServerUri
		cfg.Server.Port = envPort
		cfg.MongoDB.URI = envMongoURI
		cfg.MongoDB.Name = envDbName
		cfg.JWT.Secret = envJwtSecret
		cfg.JWT.ExpirationHours = jwtExpiry
		return &cfg, nil
	}
	
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil

}
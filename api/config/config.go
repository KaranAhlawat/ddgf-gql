package config

import "github.com/spf13/viper"

type Config struct {
	DbUser         string `mapstructure:"POSTGRES_USER"`
	DbPass         string `mapstructure:"POSTGRES_PASSWORD"`
	DbName         string `mapstructure:"POSTGRES_DB"`
	DbHost         string `mapstructure:"DB_HOST"`
	DbPort         int    `mapstructure:"DB_PORT"`
	RedisHost      string `mapstructure:"REDIS_HOST"`
	RedisPort      int    `mapstructure:"REDIS_PORT"`
	RedistPassword string `mapstructure:"REDIS_PASSWORD"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
  viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

package config

import "github.com/spf13/viper"
import "fmt"

type Config struct {
	Port          string `mapstructure:"PORT"`
	AuthSvcUrl    string `mapstructure:"AUTH_SVC_URL"`
	SessionSvcUrl string `mapstructure:"SESSION_SVC_URL"`
	RabbitMqUrl   string `mapstructure:"RABBITMQ_URL"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		fmt.Println(fmt.Errorf("Fatal error config file: %s \n", err))
		return
	}
	fmt.Println(c)

	err = viper.Unmarshal(&c)
	if err != nil {
		fmt.Println(fmt.Errorf("Fatal error config file: %s \n", err))
		return
	}

	return c, nil
}

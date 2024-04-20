package config

import "github.com/spf13/viper"

type Config struct {
	Port                 string `mapstructure:"PORT"`
	Couchbase_Connection string `mapstructure:"COUCHBASE_CONNECTION"`
	Couchbase_BucketName string `mapstructure:"COUCHBASE_BUCKETNAME"`
	Couchbase_Username   string `mapstructure:"COUCHBASE_USERNAME"`
	Couchbase_Password   string `mapstructure:"COUCHBASE_PASSWORD"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}

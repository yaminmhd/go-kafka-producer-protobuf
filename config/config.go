package config

import (
	"github.com/spf13/viper"
	"log"
)

type config struct {
	brokers []string
	topic   string
}

var appConfig config

func Load() {
	viper.AutomaticEnv()
	viper.SetConfigName("application")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()

	appConfig = config{
		brokers: getStringArray("BROKERS", true),
		topic:   getString("TOPIC", true),
	}
}

func Brokers() []string {
	return appConfig.brokers
}

func Topic() string {
	return appConfig.topic
}

func getStringArray(key string, required bool) []string{
	if required{
		checkKey(key)
	}
	return viper.GetStringSlice(key)
}

func getString(key string, required bool) string {
	if required {
		checkKey(key)
	}
	return viper.GetString(key)
}

func checkKey(key string) {
	if !viper.IsSet(key) {
		log.Panicf("Missing key: %s", key)
	}
}

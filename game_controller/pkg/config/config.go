package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port              string `mapstructure:"PORT"`
	AuthSvcUrl        string `mapstructure:"AUTH_SVC_URL"`
	ProductSvcUrl     string `mapstructure:"PRODUCT_SVC_URL"`
	OrderSvcUrl       string `mapstructure:"ORDER_SVC_URL"`
	ShipManagerSvcUrl string `mapstructure:"SHIP_MANAGER_SVC_URL"`

	PingSvcUrl string `mapstructure:"PING_SVC_URL"`
	PongSvcUrl string `mapstructure:"PONG_SVC_URL"`
}

func LoadConfig() (c Config, err error) {

	log.Println("===> [Config]: Loading config... <===")

	viper.AddConfigPath("./pks/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}

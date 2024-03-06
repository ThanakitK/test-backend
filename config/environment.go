package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func NewAppInitEnvironment() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	// viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Default --------------------------
	viper.SetDefault("app.port", 3000)
	// -----------------------------------

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Load From Environment")
	}
}

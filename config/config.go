package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

func Init() {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viperDefault()
	logConfig()
}

func viperDefault() {
	viper.SetDefault("jwt.SECRET_KEY", "your_secret_key_here")
	viper.SetDefault("jwt.TOKEN_TTL_HOURS", 24)

	viper.SetDefault("db.PATH", "./db/dish_hub.db")

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./config")

}

func logConfig() {
	log.Printf(
		"Configuration: db.PATH=%s,",
		viper.GetString("db.PATH"),
	)
}

package main

import (
	"log"

	"github.com/bulbosaur/dish-hub/config"
	"github.com/bulbosaur/dish-hub/internal/dao"
	"github.com/spf13/viper"

	_ "modernc.org/sqlite"
)

func main() {
	config.Init()

	_, err := dao.InitDB(viper.GetString("db.PATH"))
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}

}

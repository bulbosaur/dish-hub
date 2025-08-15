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

	db, err := dao.InitDB(viper.GetString("db.PATH"))
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}

	defer db.Close()

	_ = dao.NewDishHubModel(db)

	// newUser := domain.User{
	// 	Login:        "admin1",
	// 	PasswordHash: "admin",
	// }

	// uId, err := Repo.CreateUser(newUser)
	// if err != nil {
	// 	log.Fatalf("failed to create user: %v", err)
	// }

	// newRes := domain.Recipe{
	// 	Name:         "омлет",
	// 	TimeMinute:   5,
	// 	Difficulty:   "easy",
	// 	Ingredients:  []domain.Ingredient{{Name: "яйца"}},
	// 	Categories:   []string{"завтрак"},
	// 	Instructions: "разбить яйца в сковороду",
	// 	UserID:       uId,
	// }

	// id, err := Repo.Insert(newRes)
	// if err != nil {
	// 	log.Fatalf("failed to insert recipe: %v", err)
	// }

	// log.Printf("Inserted recipe with ID: %d", id)
}

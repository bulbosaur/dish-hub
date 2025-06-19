package main

import (
	"context"
	"log"

	"github.com/bulbosaur/dish-hub/config"
	"github.com/bulbosaur/dish-hub/internal/dao"
	"github.com/bulbosaur/dish-hub/internal/domain"
)

func main() {
	cfg := config.Load()

	if err := dao.InitDB(cfg); err != nil {
		log.Fatalf("failed to init db: %v", err)
	}
	defer dao.CloseDB()

	recipeDao := dao.NewRecipeDAO("recipes")

	id, err := recipeDao.Create(
		context.Background(),
		&domain.Recipe{
			Title:        "Омлет",
			PrepTime:     10,
			Difficulty:   "easy",
			Ingredients:  []domain.Ingredient{{"Яйцо", "2"}, {"молоко", "100 мл"}, {"соль", "по вкусу"}},
			Categories:   []string{"завтрак"},
			Instructions: "Разбейте яйцо в миску, влейте молоко, перемешайте, вылейте на сковороду",
		})
	if err != nil {
		log.Fatalf("failed to create recipe: %v", err)
	}
	log.Printf("created recipe with id: %s", id)

	// objectID, err := primitive.ObjectIDFromHex("68542f6a6fe61510c66d46d2")
	// if err != nil {
	// 	log.Fatalf("failed to parse id: %v", err)
	// }
	recipe, err := recipeDao.GetByID(context.Background(), id)
	if err != nil {
		log.Fatalf("failed to get recipe: %v", err)
	}
	log.Printf("recipe: %+v", recipe)
}

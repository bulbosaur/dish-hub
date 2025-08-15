package dao

import (
	"github.com/bulbosaur/dish-hub/internal/domain"
	_ "modernc.org/sqlite"
)

func (d *DishHubModel) Insert(recipe domain.Recipe) (int, error) {
	query := `INSERT INTO recipes (user_id, name, time_minute, difficulty, instructions, image_url) VALUES (?, ?, ?, ?, ?, ?)`

	result, err := d.DB.Exec(query, recipe.UserID, recipe.Name, recipe.TimeMinute, recipe.Difficulty, recipe.Instructions, recipe.ImageURL)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

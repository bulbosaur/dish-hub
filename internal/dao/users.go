package dao

import (
	"time"

	"github.com/bulbosaur/dish-hub/internal/domain"
	_ "modernc.org/sqlite"
)

func (d *DishHubModel) CreateUser(user domain.User) (int, error) {
	query := `INSERT INTO users (login, password_hash, created_at) VALUES (?, ?, ?)`

	result, err := d.DB.Exec(query, user.Login, user.PasswordHash, time.Now())
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

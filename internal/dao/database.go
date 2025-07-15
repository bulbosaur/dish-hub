package dao

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
)

// DishHubModel обрабатывает пул подключений к базе данных
type DishHubModel struct {
	DB *sql.DB
	Mu sync.Mutex
}

// NewDishHbModel создает подключение к переданной базе данных
func NewDishHubModel(db *sql.DB) *DishHubModel {
	return &DishHubModel{DB: db}
}

// InitDB открывает соединение с базой и создаёт необходимые таблицы
func InitDB(path string) (*sql.DB, error) {
	dir := filepath.Dir(path)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, fmt.Errorf("failed to create directory %s: %v", dir, err)
		}
		log.Printf("Created directory: %s", dir)
	}

	log.Printf("Database path: %s", path)

	dsn := fmt.Sprintf("file:%s?_journal_mode=WAL&_sync=NORMAL&_foreign_keys=ON", path)
	db, err := sql.Open("sqlite", dsn)

	if err != nil {
		return nil, fmt.Errorf("error when opening database: %v", err)
	}

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(0)

	createRecipes := `
 	CREATE TABLE IF NOT EXISTS recipes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		name TEXT NOT NULL,
		time_minute INTEGER,
		difficulty TEXT,
		instructions TEXT,
		image_url TEXT,
		FOREIGN KEY(user_id) REFERENCES users(id)
 	);`
	_, err = db.Exec(createRecipes)
	if err != nil {
		return nil, fmt.Errorf("error creating recipes table: %v", err)
	}

	createIngredients := `
	CREATE TABLE IF NOT EXISTS ingredients (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE
	);`
	_, err = db.Exec(createIngredients)
	if err != nil {
		return nil, fmt.Errorf("error creating ingredients table: %v", err)
	}

	createConnection := `
	CREATE TABLE IF NOT EXISTS connections (
		recipe_id INTEGER,
		ingredient_id INTEGER,
		amount TEXT,
		unit TEXT,
		PRIMARY KEY (recipe_id, ingredient_id),
		FOREIGN KEY(recipe_id) REFERENCES recipes(id),
		FOREIGN KEY(ingredient_id) REFERENCES ingredients(id)
	);`
	_, err = db.Exec(createConnection)
	if err != nil {
		return nil, fmt.Errorf("error creating connections table: %v", err)
	}

	createUsers := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		login TEXT NOT NULL UNIQUE,
		password_hash TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	_, err = db.Exec(createUsers)
	if err != nil {
		return nil, fmt.Errorf("error creating users table: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error when connecting with database: %v", err)
	}

	log.Print("Successful connection to the database")
	return db, nil
}

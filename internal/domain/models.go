package domain

type Recipe struct {
	ID           int          `json:"id"`
	Name         string       `json:"title"`
	TimeMinute   int          `json:"time_minute"`
	Difficulty   string       `json:"difficulty"`
	Ingredients  []Ingredient `json:"ingredients"`
	Categories   []string     `json:"categories"`
	Instructions string       `json:"instructions"`
	ImageURL     string       `json:"image_url"`
	UserID       int          `json:"user_id"`
}

type Connection struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Amount string `json:"amount,omitempty"`
	Unit   string `json:"unit,omitempty"`
}

type Ingredient struct {
	Name string `bson:"name" json:"name"`
}

type User struct {
	ID           int    `json:"id"`
	Login        string `json:"login"`
	PasswordHash string `json:"password"`
}

package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Recipe struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title        string             `bson:"title" json:"title"`
	PrepTime     int                `bson:"prep_time" json:"prep_time"`
	Difficulty   string             `bson:"difficulty" json:"difficulty"`
	Ingredients  []Ingredient       `bson:"ingredients" json:"ingredients"`
	Categories   []string           `bson:"categories" json:"categories"`
	Instructions string             `bson:"instructions" json:"instructions"`
	ImageURL     string             `bson:"image_url,omitempty" json:"image_url"`
	UserID       primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`
}

type Ingredient struct {
	Name   string `bson:"name" json:"name"`
	Amount string `bson:"amount" json:"amount"`
}

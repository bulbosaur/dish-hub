package dao

import (
	"context"
	"log"

	"github.com/bulbosaur/dish-hub/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecipeDAO struct {
	collection *mongo.Collection
}

func NewRecipeDAO(collectionName string) *RecipeDAO {
	return &RecipeDAO{collection: DB.Collection(collectionName)}
}

func (r *RecipeDAO) Create(ctx context.Context, recipe *domain.Recipe) (primitive.ObjectID, error) {
	result, err := r.collection.InsertOne(ctx, recipe)
	if err != nil {
		log.Printf("Error creating recipe: %v", err)
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func (r *RecipeDAO) GetByID(ctx context.Context, id primitive.ObjectID) (*domain.Recipe, error) {
	result := r.collection.FindOne(ctx, bson.M{"_id": id})

	var recipe domain.Recipe
	if err := result.Decode(&recipe); err != nil {
		log.Printf("Error decoding recipe: %v", err)
		return nil, err
	}

	return &recipe, nil
}

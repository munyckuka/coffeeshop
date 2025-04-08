package repository

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"inventory-service/domain"
)

type mongoProductRepository struct {
	collection *mongo.Collection
}

func NewMongoProductRepository(db *mongo.Database) domain.ProductRepository {
	return &mongoProductRepository{
		collection: db.Collection("products"),
	}
}

func (r *mongoProductRepository) Create(product *domain.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, product)
	return err
}

func (r *mongoProductRepository) GetByID(id string) (*domain.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var product domain.Product
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *mongoProductRepository) Update(product *domain.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(product.ID)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"name":     product.Name,
			"category": product.Category,
			"price":    product.Price,
			"stock":    product.Stock,
		},
	}

	result, err := r.collection.UpdateByID(ctx, objID, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("product not found")
	}

	return nil
}

func (r *mongoProductRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

func (r *mongoProductRepository) List(filter map[string]interface{}) ([]*domain.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []*domain.Product
	for cursor.Next(ctx) {
		var p domain.Product
		if err := cursor.Decode(&p); err != nil {
			return nil, err
		}
		products = append(products, &p)
	}

	return products, nil
}

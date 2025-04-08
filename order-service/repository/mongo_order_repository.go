package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"order-servive/domain"
)

type mongoOrderRepository struct {
	collection *mongo.Collection
}

func NewMongoOrderRepository(db *mongo.Database) domain.OrderRepository {
	return &mongoOrderRepository{
		collection: db.Collection("orders"),
	}
}

func (r *mongoOrderRepository) Create(order *domain.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	order.CreatedAt = time.Now().Unix()
	_, err := r.collection.InsertOne(ctx, order)
	return err
}

func (r *mongoOrderRepository) GetByID(id string) (*domain.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var order domain.Order
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&order)
	return &order, err
}

func (r *mongoOrderRepository) UpdateStatus(id string, status string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.UpdateOne(
		ctx,
		bson.M{"_id": objID},
		bson.M{"$set": bson.M{"status": status}},
	)

	return err
}

func (r *mongoOrderRepository) ListByUser(userID string) ([]*domain.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []*domain.Order
	for cursor.Next(ctx) {
		var o domain.Order
		if err := cursor.Decode(&o); err != nil {
			return nil, err
		}
		orders = append(orders, &o)
	}

	return orders, nil
}

package domain

type Order struct {
	ID        string      `json:"id" bson:"_id,omitempty"`
	UserID    string      `json:"user_id" bson:"user_id"`
	Items     []OrderItem `json:"items" bson:"items"`
	Status    string      `json:"status" bson:"status"` // pending, completed, cancelled
	CreatedAt int64       `json:"created_at" bson:"created_at"`
}

type OrderItem struct {
	ProductID string `json:"product_id" bson:"product_id"`
	Quantity  int    `json:"quantity" bson:"quantity"`
}

type OrderRepository interface {
	Create(order *Order) error
	GetByID(id string) (*Order, error)
	UpdateStatus(id string, status string) error
	ListByUser(userID string) ([]*Order, error)
}

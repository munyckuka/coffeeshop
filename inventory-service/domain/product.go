package domain

type Product struct {
	ID       string  `json:"id" bson:"_id,omitempty"`
	Name     string  `json:"name" bson:"name"`
	Category string  `json:"category" bson:"category"`
	Price    float64 `json:"price" bson:"price"`
	Stock    int     `json:"stock" bson:"stock"`
}

type ProductRepository interface {
	Create(product *Product) error
	GetByID(id string) (*Product, error)
	Update(product *Product) error
	Delete(id string) error
	List(filter map[string]interface{}) ([]*Product, error)
}

package entity

import "github.com/google/uuid"

type Category struct {
	ID   string
	Name string
}

func NewCategory(name string) *Category {
	return &Category{
		ID:   uuid.New().String(),
		Name: name,
	}
}

type Product struct {
	ID string
	Name string
	Description string
	Price float64
	CategoryID string
	ImageURL string
}

func NewProduct(name string, descriptrion string, price float64, categoryID string, ImageURL string) *Product {
	return &Product{
		ID: uuid.New().String(),
		Name: name,
		Description: descriptrion,
	  Price: price,
		CategoryID: categoryID,
		ImageURL: ImageURL,
	}
}
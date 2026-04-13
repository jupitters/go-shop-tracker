package models

import "github.com/jinzhu/gorm"

var (
	OrderStatus = []string{"Order placed", "Preparing", "Baking", "Quality check", "Ready"}

	PizzaTypes = []string{
		"Margherita",
		"Pepperoni",
		"Vegetarian",
		"Hawaiian",
		"Bbq Chiken",
		"Meat Lovers",
		"Buffalo Chiken",
		"Supreme",
		"Truffle Mushroom",
		"Four Cheese",
	}

	PizzaSizes = []string{
		"Small", "Medium", "Large", "X-Large",
	}
)

type OrderModel struct {
	DB *gorm.DB
}

type Order struct {
	ID     string `gorm:"primaryKey;size:14" json:"id"`
	Status string `gorm:"not null" json:"status"`
}

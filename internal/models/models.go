package models

import (
	"fmt"

	"gorm.io/gorm"
	"modernc.org/sqlite"
)

type DBModel struct {
	OrderModel
}

func InitDB(dataSourceName string) (*DBModel, error) {
	db, err := gorm.Open(sqlite.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("Failed to migrate database: %v", err)
	}

}

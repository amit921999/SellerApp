package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Id          string  `json:"id" gorm:"size:64"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Quantity    int64   `json:"quantity"`
	OrderId     string  `gorm:"size:64"`
}

type ItemResponse struct {
	Id          string  `json:"id" gorm:"size:64"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Quantity    int64   `json:"quantity"`
}

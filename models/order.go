package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Id           string  `json:"id" gorm:"size:64"`
	Status       string  `json:"status"`
	Total        float32 `json:"total"`
	CurrencyUnit string  `json:"currencyUnit"`
	Items        []Item  `json:"items"`
}
type CreateOrderResponse struct {
	OrderResponse []OrderResponse
}

type OrderResponse struct {
	Id           string         `json:"id" gorm:"size:64"`
	Status       string         `json:"status"`
	Total        float32        `json:"total"`
	CurrencyUnit string         `json:"currencyUnit"`
	Items        []ItemResponse `json:"items"`
}

type UpdateOrder struct {
	Status       string  `json:"status"`
	Total        float32 `json:"total"`
	CurrencyUnit string  `json:"currencyUnit"`
}

func Appending(odr []Order) []OrderResponse {
	var orders []OrderResponse
	for _, order := range odr {
		itemsRes := CovertItem(order)
		orders = append(orders, OrderResponse{
			Id:           order.Id,
			Status:       order.Status,
			Total:        order.Total,
			CurrencyUnit: order.CurrencyUnit,
			Items:        itemsRes,
		})
	}
	return orders
}

func CovertItem(order Order) []ItemResponse {
	var itemsRes []ItemResponse
	for _, item := range order.Items {
		itemsRes = append(itemsRes, ItemResponse{
			Id:          item.Id,
			Description: item.Description,
			Price:       item.Price,
			Quantity:    item.Quantity,
		})
	}
	return itemsRes
}

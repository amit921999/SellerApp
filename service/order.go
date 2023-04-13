package service

import (
	"orderManagement/initialize"
	"orderManagement/models"
)

func CreateOrder(order models.Order) models.OrderResponse {
	result := initialize.DB.Create(&order)
	if result.Error != nil {
		panic(result.Error)
	}
	return models.OrderResponse{
		Id:           order.Id,
		Status:       order.Status,
		Total:        order.Total,
		CurrencyUnit: order.CurrencyUnit,
		Items:        models.CovertItem(order),
	}
}

func GetOrder(order models.FilterOrd) models.CreateOrderResponse {
	var odr []models.Order
	ords := &models.Order{Id: order.Id, Status: order.Status, Total: order.Total, CurrencyUnit: order.CurrencyUnit}

	result := initialize.DB.Preload("Items").Find(&odr, ords)
	if result.Error != nil {
		panic(result.Error)
	}

	orders := models.Appending(odr)

	return models.CreateOrderResponse{
		OrderResponse: orders,
	}
}

func UpdateOrder(order models.UpdateOrder, orderId string) models.OrderResponse {
	var odr models.Order

	odr.Id = orderId
	odr.Status = order.Status
	odr.CurrencyUnit = order.CurrencyUnit
	odr.Total = order.Total

	result := initialize.DB.Model(&odr).Updates(odr)
	if result.Error != nil || result.RowsAffected == 0 {
		panic(result.Error)
	}

	result = initialize.DB.Preload("Items").First(&odr, "id=?", orderId)
	if result.Error != nil || result.RowsAffected == 0 {
		panic(result.Error)
	}

	return models.OrderResponse{
		Id:           odr.Id,
		Status:       odr.Status,
		Total:        odr.Total,
		CurrencyUnit: odr.CurrencyUnit,
		Items:        models.CovertItem(odr),
	}
}

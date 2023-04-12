package models

type FilterOrd struct {
	Id           string  `form:"id" gorm:"size:64"`
	Status       string  `form:"status"`
	Total        float32 `form:"total"`
	CurrencyUnit string  `form:"currencyUnit"`
}

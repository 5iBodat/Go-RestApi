package models

type Products struct {
	Id    int     `json:"id" gorm:"primaryKey"`
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

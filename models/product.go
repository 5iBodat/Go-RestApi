package models

type Products struct {
	Id    int     `json:"id" gorm:"default:1"`
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

package model

type Product struct {
	ID    int64  `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

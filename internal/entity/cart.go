package entity

import "time"

type Cart struct {
	Id        string
	UserId    string // foreign key to user table
	ProductId string // foreign key to product table
	Quantity  int
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt *time.Time
	UpdatedBy *string

	Product *Product // reference to product table, default is nil
}

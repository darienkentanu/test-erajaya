package models

import "time"

type Product struct {
	ID          int    `gorm:"primarykey;AUTO_INCREMENT" json:"id" form:"id"`
	Name        string `gorm:"type:varchar(255);unique;not null" json:"name" form:"name"`
	Price       int    `gorm:"type:int;not null" json:"price" form:"price"`
	Description string `gorm:"type:longtext;not null" json:"description" form:"description"`
	Quantity    int    `gorm:"type:int;not null" json:"quantity" form:"quantity"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

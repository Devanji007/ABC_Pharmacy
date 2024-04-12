package models

import (
	"time"

)

type Items struct {
	ItemId    	uint     	`gorm:"primary key;autoIncrement" json:"id"`
	ItemName  	*string  	`json:"itemname"`
	UnitPrice 	*float64 	`json:"unitprice"`
	Category  	*string  	`json:"category"`
	CreatedAt 	time.Time 	`json:"createdat"`
	UpdatedAt   time.Time	`json:"updatedat"`
}


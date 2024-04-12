package models

import (
	"time"

	
)

type Invoice struct {
	InvId       uint      `gorm:"primary key;autoIncrement" json:"invid"`
	CusName     *string   `json:"cusname"`
	MobNo       *string   `json:"mobno"`
	Email       *string   `json:"email"`
	Address     *string   `json:"address"`
	BillingType *string   `json:"billingtype"`
	TotalAmount *float64  `json:"totalamount"`
	CreatedAt   time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}



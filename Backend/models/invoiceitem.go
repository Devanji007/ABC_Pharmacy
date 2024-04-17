package models

type InvoiceItem struct {
	InvId     uint   `gorm:"primaryKey" json:"invid"`
	ItemId    uint   `gorm:"primaryKey" json:"itemid"`
	Quantity  int64  `json:"quantity"`
	UnitPrice float64 `json:"unitprice"`
}

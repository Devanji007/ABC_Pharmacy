package routes

import (
	"ABC_PHARMACY/controllers"
	"time"

	"github.com/gofiber/fiber/v2"
)

type InvoiceItem struct {
	InvId     uint   `gorm:"primaryKey" json:"invid"`
	ItemId    uint   `gorm:"primaryKey" json:"itemid"`
	Quantity  int64  `json:"quantity"`
	UnitPrice float64 `json:"unitprice"`
}

type Invoice struct {
	InvId       uint      `gorm:"primary key;autoIncrement" json:"invid"`
	CusName     *string   `json:"cusname"`
	MobNo       *string   `json:"mobno"`
	Email       *string   `json:"email"`
	Address     *string   `json:"address"`
	BillingType *string   `json:"billingtype"`
	TotalAmount *string   `json:"totalamount"`
	CreatedAt   time.Time `json:"createdat"`
	UpdatedAt   time.Time `json:"updatedat"`
}

func SetupInvoiceRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_invoice", controllers.CreateInvoice)
	//api.Put("/edit_item", controllers.UpdateItem)
	api.Delete("delete_invoice/:inv_id", controllers.DeleteInvoice)
	api.Get("/get_invoices/:inv_id", controllers.GetInvoiceByID)
	api.Get("/invoices", controllers.GetInvoices)
}

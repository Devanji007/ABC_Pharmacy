package routes

import (
	"ABC_PHARMACY/controllers"
	"time"

	"github.com/gofiber/fiber/v2"
)

type InvoiceItem struct {
	ItemName  string  `json:"itemname"`
	UnitPrice float64 `json:"unitprice"`
	Category  string  `json:"category"`
}

type Invoice struct {
	InvId       uint      `gorm:"primary key;autoIncrement" json:"invid"`
	CusName     *string   `json:"cusname"`
	MobNo       *string   `json:"mobno"`
	Email       *string   `json:"email"`
	Address     *string   `json:"address"`
	BillingType *string   `json:"billingtype"`
	TotalAmount *float64  `json:"totalamount"`
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

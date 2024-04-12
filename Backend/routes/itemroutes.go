package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"ABC_PHARMACY/controllers"
)

type Item struct {
	ItemName  string    `json:"itemname"`
	UnitPrice float64   `json:"unitprice"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}

func SetupItemRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/add_items", controllers.AddItem)
	api.Put("/edit_item", controllers.UpdateItem)
	api.Delete("delete_item/:item_id", controllers.DeleteItem)
	api.Get("/get_items/:item_id", controllers.GetItemByID)
	api.Get("/items", controllers.GetItems)
}

package controllers

import (
	"ABC_PHARMACY/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Item struct {
	ItemId    uint      `gorm:"primary key;autoIncrement" json:"id"`
	ItemName  string    `json:"itemname"`
	UnitPrice float64   `json:"unitprice"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}

func AddItem(context *fiber.Ctx) error {
	item := Item{}

	err := context.BodyParser(&item)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"meesage": "request failed"})
		return err
	}

	err = r.DB.Create(&item).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not add item"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "item has been added"})
	return nil
}

func UpdateItem(context *fiber.Ctx) error {
	// Parse request body into an Item struct
	var updateData models.Items
	if err := context.BodyParser(&updateData); err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Invalid request body",
		})
		return err
	}

	// Get item ID from the URL parameter
	id := updateData.ItemId

	// Get the existing item from the database
	var existingItem models.Items
	if err := r.DB.Where("item_id = ?", id).First(&existingItem).Error; err != nil {
		context.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "Item not found",
		})
		return err
	}

	// Update the existing item with values from the request body
	existingItem.ItemName = updateData.ItemName
	existingItem.UnitPrice = updateData.UnitPrice
	existingItem.Category = updateData.Category

	// Set the UpdatedAt field to the current time
	existingItem.UpdatedAt = time.Now()

	// Save the updated item to the database
	if err := r.DB.Where("item_id = ?", id).Updates(&existingItem).Error; err != nil {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "Failed to update item",
		})
		return err
	}

	// Format the UpdatedAt time

	updatedAtStr := existingItem.UpdatedAt.Format("15:04:05")

	// Format the CreatedAt time

	createdAtStr := existingItem.CreatedAt.Format("15:04:05")

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message":    "Item updated successfully",
		"data":       existingItem,
		"updated_at": updatedAtStr,
		"created_at": createdAtStr,
	})
	return nil
}

func DeleteItem(context *fiber.Ctx) error {

	id := context.Params("item_id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil

	}

	// Delete the item from the database using its ID
	if err := r.DB.Where("item_id = ?", id).Delete(&models.Items{}).Error; err != nil {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "Failed to delete item",
		})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "item deleted successfully",
	})
	return nil
}

func GetItems(context *fiber.Ctx) error {
	itemModels := &[]models.Items{}

	err := r.DB.Find(itemModels).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get items"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "items fetch successfully",
		"data":    itemModels,
	})

	return nil
}

func GetItemByID(context *fiber.Ctx) error {

	id := context.Params("item_id")
	itemModel := &models.Items{}
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	fmt.Println("the Id is", id)
	err := r.DB.Where("item_id = ?", id).First(itemModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"meesage": "could not get the item"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "item id fetched succcessfully",
		"data":    itemModel,
	})
	return nil
}

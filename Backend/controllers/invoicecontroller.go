package controllers

import (
	"ABC_PHARMACY/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type InvoiceItem struct {
	InvId    uint  `gorm:"primaryKey" json:"invid"`
	ItemId   uint  `gorm:"primaryKey" json:"itemid"`
	Quantity int64 `json:"unitprice"`
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

func CreateInvoice(context *fiber.Ctx) error {
	invoice := Invoice{}

	err := context.BodyParser(&invoice)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"meesage": "request failed"})
		return err
	}

	err = r.DB.Create(&invoice).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create invoice"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "invoice has been created"})
	return nil
}

func DeleteInvoice(context *fiber.Ctx) error {

	id := context.Params("inv_id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil

	}

	// Delete the invoice from the database using its ID
	if err := r.DB.Where("inv_id = ?", id).Delete(&models.Invoice{}).Error; err != nil {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "Failed to delete invoice",
		})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "invoice deleted successfully",
	})
	return nil
}

func GetInvoices(context *fiber.Ctx) error {
	invoiceModels := &[]models.Invoice{}

	err := r.DB.Find(invoiceModels).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get invoices"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "invoices fetch successfully",
		"data":    invoiceModels,
	})

	return nil
}

func GetInvoiceByID(context *fiber.Ctx) error {

	id := context.Params("inv_id")
	invoiceModel := &models.Invoice{}
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	fmt.Println("the Id is", id)
	err := r.DB.Where("inv_id = ?", id).First(invoiceModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"meesage": "could not get the invoice"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "invoice id fetched succcessfully",
		"data":    invoiceModel,
	})
	return nil
}

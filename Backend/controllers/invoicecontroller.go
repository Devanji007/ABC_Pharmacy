package controllers

import (
	"ABC_PHARMACY/models"
	"fmt"
	"net/http"
	

	"github.com/gofiber/fiber/v2"
)

type InvoiceReq struct{
	Invoice models.Invoice 		`json:"invoice"`
	Items []models.InvoiceItem 	`json:"items"`
}

func CreateInvoice(context *fiber.Ctx) error {
	invoiceReq := InvoiceReq{}

	err := context.BodyParser(&invoiceReq)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"meesage": "request failed"})
		return err
	}

	invoice := invoiceReq.Invoice
	items := invoiceReq.Items

	result := r.DB.Create(&invoice)
	if result.Error != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create invoice"})
		return err
	}

	// Iterate over the items in the Invoice struct
	for _, item := range items {
		// Create a new InvoiceItem record
		invoiceItem := models.InvoiceItem{
			InvId:     invoice.InvId, // Assign the InvId from the Invoice struct
			ItemId:    item.ItemId,   // Assign the ItemId from the Item struct
			Quantity:  item.Quantity,
			UnitPrice: item.UnitPrice,
		}

		// Insert the new record into the InvoiceItem table
		if err := r.DB.Create(&invoiceItem).Error; err != nil {
			context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "Failed to create invoice item"})
			return err
		}
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
	invoiceItemModel := []models.InvoiceItem{}

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

	err = r.DB.Where("inv_id = ?", id).Find(&invoiceItemModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"meesage": "could not get the invoice"})
		return err
	}

	invoice := InvoiceReq{
		Invoice: *invoiceModel,
		Items: invoiceItemModel,
	}


	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "invoice id fetched succcessfully",
		"data":    invoice,
	})
	return nil
}

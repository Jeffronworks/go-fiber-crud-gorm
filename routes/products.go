package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/jeffronworks/fiber-api/database"
	"github.com/jeffronworks/fiber-api/models"
)

type Product struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func CreateResponseProduct(productModel models.Product) Product {
	return Product{ID: productModel.ID, Name: productModel.Name, SerialNumber: productModel.SerialNumber}
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&product)

	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}

func GetProducts(c *fiber.Ctx) error {
	products := []models.Product{}

	database.Database.Db.Find(&products)
	responseProducts := []Product{}

	for _, product := range products {
		responseProduct := CreateResponseProduct(product)
		responseProducts = append(responseProducts, responseProduct)
	}

	return c.Status(200).JSON(responseProducts)

}

func findProduct(id int, product *models.Product) error {
	database.Database.Db.Find(&product, "id = ?", id) // find the user in the database

	if product.ID == 0 {
		return errors.New("user does not exist")
	}

	return nil
}

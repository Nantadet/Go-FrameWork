package handler

import (
	"errors"
	"strconv"

	"github.com/Nantadet/go-fiber-test/model"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type productHandler struct {
	db *gorm.DB
}

func NewSchoolHandler(db *gorm.DB) *productHandler {
	return &productHandler{db: db}
}

func (h *productHandler) GetProducts(c fiber.Ctx) error {
	var products []model.Product
	if err := h.db.Find(&products).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch products",
		})
	}

	return c.JSON(products)

}

func (h *productHandler) CreateProduct(c fiber.Ctx) error {

	var product model.Product

	if err := c.Bind().Body(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.db.Create(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create product",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(product)

}

func (h *productHandler) GetProduct(c fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid product id",
		})
	}

	var product model.Product
	if err := h.db.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "product not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch product",
		})
	}

	return c.JSON(product)
}

func (h productHandler) DeleteProductByID(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}
	var product []model.Product
	if err := h.db.Delete(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "product not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch product",
		})
	}
	return c.JSON("delated")

}

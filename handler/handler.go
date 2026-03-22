package handler

import (
	"errors"
	"strconv"
	"strings"

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

func (h *productHandler) UpdateProductByID(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalit request body.",
		})
	}
	var req model.Product
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalit request body",
		})
	}
	req.Name = strings.TrimSpace(req.Name)
	Pric := req.Price
	if req.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "request Name",
		})
	}
	var products model.Product
	if err := h.db.First(&products, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Product not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error fetching school",
		})
	}
	products.Name = req.Name
	products.Price = Pric

	if err := h.db.Save(&products).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error updating schools",
		})
	}
	return c.Status(fiber.StatusOK).JSON(products)
}

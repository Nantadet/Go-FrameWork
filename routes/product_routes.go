package routes

import (
	"github.com/Nantadet/go-fiber-test/handler"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	productHandler := handler.NewSchoolHandler(db)
	api := app.Group("/api")

	api.Get("/product", productHandler.GetProducts)
	api.Get("/product/:id", productHandler.GetProduct)
	api.Post("/product", productHandler.CreateProduct)
	api.Delete("/product/:id", productHandler.DeleteProductByID)
	api.Put("/product/:id", productHandler.UpdateProductByID)
}

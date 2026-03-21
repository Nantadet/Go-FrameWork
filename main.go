package main

import (
	"fmt"
	"log"

	"github.com/Nantadet/go-fiber-test/model"
	"github.com/Nantadet/go-fiber-test/routes"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/static"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = "3306"
	database = "mydb"
	user     = "root"
	password = ""
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	db.AutoMigrate(&model.Product{})
	fmt.Println("connect success", db != nil)

	app := fiber.New(fiber.Config{
		AppName: "HLLC 2026 - Backend Test",
	})

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use("/", static.New("./public"))
	routes.SetupRoutes(app, db)

	log.Println(`
		🚀 NANTADET DEV SERVER

        ┌─────────────────────────────┐
        │ Status : Running            │
        │ Local  : http://localhost:8080
        │ Mode   : Development        │
        │ Ready  : ✔                  │
		  Server starting on :8080`)
	log.Fatal(app.Listen(":3000"))
}

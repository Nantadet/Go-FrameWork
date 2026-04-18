package main

import (
	"fmt"
	"log"

	"github.com/Nantadet/go-fiber-test/config"
	"github.com/Nantadet/go-fiber-test/model"
	"github.com/Nantadet/go-fiber-test/routes"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/static"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func clearConsole() {
	fmt.Print("\033[H\033[2J")
}

const Port = 3000

func main() {
	clearConsole()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Ohlcv{})
	db.AutoMigrate(&model.Order{})
	fmt.Println("connect success", db != nil)

	app := fiber.New(fiber.Config{
		AppName: "NANTADET LEARN-GO",
	})

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use("/", static.New("./public"))
	routes.SetupRoutes(app, db)

	fmt.Printf(`
	🚀 NANTADET DEV SERVER

        ┌─────────────────────────────┐
        │ Status : Running            │
        │ Local  : http://localhost:%d │
        │ Mode   : Development        │
        │ Ready  : OK                 │
        └─────────────────────────────┘`, Port)

	log.Fatal(app.Listen(":3000", fiber.ListenConfig{
		DisableStartupMessage: true,
	}))
}

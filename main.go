package main

import (
	"fmt"

	"github.com/burabatbold/delivery-auth-service/database"
	"github.com/burabatbold/delivery-auth-service/grpc"
	configPkg "github.com/burabatbold/delivery-auth-service/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
)

func main() {

	configPkg.NewConfig()
	database.Connect()

	app := fiber.New(fiber.Config{
		BodyLimit: 1024 * 1024 * 10, // 10MB
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	go grpc.StartServer()

	app.Listen(fmt.Sprintf(":%d", viper.GetInt("app.port")))
}

package main

import (
	"fmt"

	"github.com/anjiri1684/ticket-booking-project-v1/config"
	"github.com/anjiri1684/ticket-booking-project-v1/db"
	"github.com/anjiri1684/ticket-booking-project-v1/handlers"
	"github.com/anjiri1684/ticket-booking-project-v1/middleware"
	"github.com/anjiri1684/ticket-booking-project-v1/repositories"
	"github.com/anjiri1684/ticket-booking-project-v1/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	envConfig := config.NewEnConfig()

	db := db.Init(envConfig, db.DBMigrator)

	

	app := fiber.New(fiber.Config{
		AppName: "Ticket Booking",
		ServerHeader: "Fiber",
	})

	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins:     "http://192.168.0.100:19006,exp://192.168.0.100:19000",
	// 	AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
	// 	AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
	// 	AllowCredentials: true,
	// }))
	app.Use(cors.New())

	

	//repository
	eventRepository := repositories.NewEventRepository(db)
	ticketRepository := repositories.NewTicketRepository(db)
	authRepository := repositories.NewAuthRepository(db)


	//service
	authService := services.NewAuthServices(authRepository)

	//routing
	server := app.Group("/api")
	handlers.NewAuthHanlder(server.Group("/auth"), authService)


	privateRoutes := server.Use(middleware.AuthProtected(db))


	//handlers
	handlers.NewEventHandler(privateRoutes.Group("/event"), eventRepository)
	handlers.NewTicketHandler(privateRoutes.Group("/ticket"), ticketRepository)

	//listen the server
	app.Listen(fmt.Sprintf("%s",":" + envConfig.ServerPort))
}
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
)

func main() {
	envConfig := config.NewEnConfig()

	db := db.Init(envConfig, db.DBMigrator)

	

	app := fiber.New(fiber.Config{
		AppName: "Ticket Booking",
		ServerHeader: "Fiber",
	})

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


	app.Listen(fmt.Sprintf("%s",":" + envConfig.ServerPort))
}
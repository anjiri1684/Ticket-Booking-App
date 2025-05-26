package handlers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/anjiri1684/ticket-booking-project-v1/models"
	"github.com/gofiber/fiber/v2"
	"github.com/skip2/go-qrcode"
)

type TicketHandler struct {
	repository models.TicketRepository
}

func(h *TicketHandler)GetMany(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5 * time.Second))

	defer cancel()

	userId := ctx.Locals("userId").(uint)


	tickets, err := h.repository.GetMany(context, userId)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status": "Fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"message": "Available tickets",
		"data": tickets,
	})

}

func (h *TicketHandler)GetOne(ctx *fiber.Ctx) error {
	
	ticketId , _ := strconv.Atoi(ctx.Params("ticketId"))

	
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5 * time.Second))

	defer cancel()

	userId := ctx.Locals("userId").(uint)



	ticket, err := h.repository.GetOne(context,userId, uint(ticketId))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "fail",
			"message": "No tickets found",

		})
	}

	var QRCode []byte

	QRCode, err = qrcode.Encode(
		fmt.Sprintf("tocketId:%v ownerId:%v", ticketId, userId),
		qrcode.Medium,
		256,
	)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "Fail",
			"message": "Failed tovalidate",
			"error": ticket,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"message": "Generated QRcode",
		"data": fiber.Map{
			"ticket": ticket,
			"qrcode": QRCode,
		},
	})
}


func (h *TicketHandler) CreateOne(ctx *fiber.Ctx) error {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	newTicket := &models.Ticket{}
	if err := ctx.BodyParser(&newTicket); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status":  "fail",
			"message": "Invalid payload",
		})
	}

	// Validate foreign key reference
	// var event models.Event
	// if err := h.repository.DB().First(&event, newTicket.EventID).Error; err != nil {
	// 	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"Status":  "fail",
	// 		"message": "Event with the provided ID does not exist",
	// 	})
	// }

	userId := ctx.Locals("userId").(uint)


	ticket, err := h.repository.CreateOne(ctxWithTimeout,userId, newTicket)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status":  "fail",
			"message": "Ticket creation failed due to DB constraints",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"Status":  "success",
		"Message": "Ticket created successfully",
		"data":    ticket,
	})
}


func (h *TicketHandler) ValidateOne(ctx *fiber.Ctx) error {
	// Context with timeout
	reqCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Parse request body
	validateBody := new(models.ValidateTicket)
	if err := ctx.BodyParser(validateBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "Failed",
			"message": "Invalid request payload",
			"error":   err.Error(),
		})
	}

	// 🚨 Validate required fields
	if validateBody.TicketId == 0 || validateBody.OwnerId == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "Failed",
			"message": "Ticket ID and Owner ID must be greater than 0",
		})
	}

	// Prepare update payload
	validateData := map[string]interface{}{
		"entered": true,
	}

	// Execute update in repo
	validTicket, err := h.repository.UpdateOne(reqCtx, validateBody.OwnerId, validateBody.TicketId, validateData)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  "Failed",
			"message": "Ticket validation failed",
			"error":   err.Error(),
		})
	}

	// Success response
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "Success",
		"message": "Ticket validation successful",
		"data":    validTicket,
	})
}


func NewTicketHandler(router fiber.Router, repository models.TicketRepository) {
	handler := &TicketHandler{
		repository: repository,
	}


	router.Get("/", handler.GetMany)
	router.Post("/", handler.CreateOne)
	router.Get("/:ticketId", handler.GetOne)
	router.Post("/validate", handler.ValidateOne)
}
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

	userId := uint(ctx.Locals("userId").(float64))

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

	userId := uint(ctx.Locals("userId").(float64))


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

	userId := uint(ctx.Locals("userId").(float64))

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


func (h *TicketHandler)ValidateOne(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5 * time.Second))
	
	defer cancel()

	validateBody := &models.ValidateTicket{}

	if err := ctx.BodyParser(&validateBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status": "Failed",
			"message": "failed to validate",
			"data": nil,
		})
	}

	validateData := make(map[string]interface{})
	validateData["entered"] = true


	validTicket, err := h.repository.UpdateOne(context,validateBody.OwnerId, validateBody.TicketId, validateData)

	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"Status": "Failed",
			"message": "Ticket validation failed",
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"message": "Ticket Validation success",
		"data": validTicket,
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
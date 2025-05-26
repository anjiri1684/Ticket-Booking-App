package handlers

import (
	"context"
	"strconv"
	"time"

	"github.com/anjiri1684/ticket-booking-project-v1/models"
	"github.com/gofiber/fiber/v2"
)

type EventHandler struct {
	repository models.EventRepository
}

func (h *EventHandler)GetMany(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5 * time.Second))

	defer cancel()

	events, err := h.repository.Getmany(context)

	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status": "Failed",
			"message" : err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Status": "Success",
		"message": "",
		"data": events,
	})
}

func (h *EventHandler)PostOne(ctx *fiber.Ctx) error {
	return nil
}

func (h *EventHandler)GetOne(ctx *fiber.Ctx) error {
	eventId, _ := strconv.Atoi(ctx.Params("eventId"))

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5 *time.Second))

	defer cancel()
	
	event, err := h.repository.GetOne(context, uint(eventId))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Status": "Success",
		"message": "Succesfully retrived event",
		"data": event,
	})

}

func (h *EventHandler) CreateOne(c *fiber.Ctx) error {
	event := &models.Event{}
 
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
 
	if err := c.BodyParser(event); err != nil {
	    return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
		   "status":  "fail",
		   "message": err.Error(),
		   "data":    nil,
	    })
	}
 
	createdEvent, err := h.repository.CreateOne(ctxTimeout, event)
	if err != nil {
	    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		   "status":  "fail",
		   "message": err.Error(),
	    })
	}
 
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
	    "status":  "success",
	    "message": "Event created successfully",
	    "data":    createdEvent,
	})
 }
 

func (h *EventHandler)UpdateOne(ctx *fiber.Ctx) error {
	eventId, _ := strconv.Atoi(ctx.Params("eventId"))

	updatedata := make(map[string]interface{})

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5 * time.Second))

	defer cancel()

	if err := ctx.BodyParser(&updatedata); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status": "Failed",
			"message" : err.Error(),
			"data": nil,
		})
	}

	event, err := h.repository.UpdateOne(context,uint(eventId), updatedata)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"Status": "Success",
		"message": "Event created",
		"data": event,
	})

}

func (h *EventHandler)DeleteOne(ctx *fiber.Ctx) error {
	eventId, _ := strconv.Atoi(ctx.Params("eventId"))

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5 * time.Second))

	defer cancel()
	err := h.repository.DeleteOne(context, uint(eventId))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func NewEventHandler(router fiber.Router, repository models.EventRepository) {
	handler := &EventHandler{
		repository: repository,
	}

	router.Get("/", handler.GetMany)
	router.Post("/", handler.CreateOne)
	router.Get("/:eventId", handler.GetOne)
	router.Put("/:eventId", handler.UpdateOne)
	router.Delete("/:eventId", handler.DeleteOne)

}
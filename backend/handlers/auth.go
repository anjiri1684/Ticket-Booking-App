package handlers

import (
	"context"
	"time"

	"github.com/anjiri1684/ticket-booking-project-v1/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type AuthHandler struct {
	service models.AuthService
}


func (h *AuthHandler) Login(ctx *fiber.Ctx) error {
	creds := &models.AuthCredentials{}
 
	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
 
	if err := ctx.BodyParser(&creds); err != nil {
	    return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		   "status":  "fail",
		   "message": "Failed to parse login request",
		   "error":   err.Error(),
	    })
	}
 
	if err := validate.Struct(creds); err != nil {
	    return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		   "status":  "fail",
		   "message": "Please provide a valid email and password",
		   "error":   err.Error(),
	    })
	}
 
	token, user, err := h.service.Login(context, creds)
	if err != nil {
	    return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		   "status":  "fail",
		   "message": "Failed to login",
		   "error":   err.Error(),
	    })
	}
 
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
	    "status":  "success",
	    "message": "login success",
	    "data": fiber.Map{
		   "user": fiber.Map{
			  "id":    user.ID,
			  "email": user.Email,
			  "role":  user.Role,
		   },
		   "token": token,
	    },
	})
 }
 

func (h *AuthHandler)Register(ctx *fiber.Ctx)error {
	creds := &models.AuthCredentials{}

	context, cancel := context.WithTimeout(context.Background(), time.Duration( 5 *time.Second))

	defer cancel()

	if err := ctx.BodyParser(&creds); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "Fail",
			"message": "Failed to register",
			"error": err.Error(),
		})
	}

	if err := validate.Struct(creds); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "Fail",
			"message": "Please provide a valid email and password",
			"error": err.Error(),
		})
	}

	token, user, err := h.service.Register(context, creds)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "Fail",
			"message": "Failed to Register",
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Status": "Success",
		"message": "login success",
		"token": token,
		"data": fiber.Map{
			"id": user.ID,
			"email": user.Email,
			"role": user.Role,
			
		},
	})

}

func NewAuthHanlder(router fiber.Router, service models.AuthService) {
	handler := &AuthHandler {
		service: service,
	}

	router.Post("/login", handler.Login)
	router.Post("/register", handler.Register)
}
package middleware

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/anjiri1684/ticket-booking-project-v1/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func AuthProtected(db *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
	    authHeader := ctx.Get("Authorization")
	    if authHeader == "" {
		   log.Warn("Missing Authorization header")
		   return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			  "status":  "fail",
			  "message": "Unauthorized",
		   })
	    }
 
	    parts := strings.SplitN(authHeader, " ", 2)
	    if len(parts) != 2 || parts[0] != "Bearer" {
		   log.Warn("Invalid Authorization format")
		   return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			  "status":  "fail",
			  "message": "Unauthorized",
		   })
	    }
 
	    tokenString := parts[1]
	    secret := []byte(os.Getenv("JWT_SECRET"))
 
	    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		   if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			  return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		   }
		   return secret, nil
	    })
 
	    if err != nil || !token.Valid {
		   log.Warn("Invalid or expired token")
		   return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			  "status":  "fail",
			  "message": "Unauthorized",
		   })
	    }
 
	    claims, ok := token.Claims.(jwt.MapClaims)
	    if !ok || !token.Valid {
		   return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			  "status":  "fail",
			  "message": "Invalid token claims",
		   })
	    }
 
	    userIDFloat, ok := claims["id"].(float64)
	    if !ok {
		   return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			  "status":  "fail",
			  "message": "Invalid token payload",
		   })
	    }
	    userID := uint(userIDFloat)
 
	    var user models.User
	    if err := db.First(&user, "id = ?", userID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		   log.Warn("User not found in DB")
		   return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			  "status":  "fail",
			  "message": "Unauthorized",
		   })
	    }
 
	    ctx.Locals("userId", userID)
	    return ctx.Next()
	}
 }
 
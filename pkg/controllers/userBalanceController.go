package controllers

import (
	"net/http"

	"github.com/faruqii/Startup-Campus-Final-Project-Backend/pkg/database"
	"github.com/faruqii/Startup-Campus-Final-Project-Backend/pkg/models"
	"github.com/gofiber/fiber/v2"
)

// User Topup Balance
func TopupBalance(c *fiber.Ctx) error {
	req := models.UserBalanceRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	user := c.Locals("user").(models.UserToken)

	balance := models.UserBalance{}

	err := database.DB.Where("user_id = ?", user.UserID).First(&balance).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	balance.Balance = balance.Balance + req.Balance

	err = database.DB.Save(&balance).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success topup balance",
	})

}

// User Get Balance
func GetBalance(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserToken)

	var userBalance models.UserBalance

	err := database.DB.Where("user_id = ?", user.UserID).First(&userBalance).Error
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"balance": userBalance.Balance,
	})
}

package handlers

import (
	services "think-intern-2023/services"

	"github.com/gofiber/fiber/v2"
)

type thinkHandler struct {
	thinkService services.ThinkService
}

func NewThinkHandler(thinkService services.ThinkService) thinkHandler {
	return thinkHandler{
		thinkService: thinkService,
	}
}
func (h thinkHandler) GetThink(c *fiber.Ctx) error {
	var request services.Request
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	response, err := h.thinkService.Think(request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(response)
}

// ไม่จำเป็นต้องใช้ service
// ไม่จำเป็นต้องใช้ service
// ไม่จำเป็นต้องใช้ service

package handler

import "github.com/gofiber/fiber/v2"

// Hello hanlde api status godoc
// @Summary     Api status
// @Description Verify api is running
// @Tags        status
// @Accept      json
// @Produce     plain
// @Success     200 {string} json   "Status equals success"
// @Failure     400 {string} string "ok"
// @Failure     404 {string} string "ok"
// @Failure     500 {string} string "ok"
// @Router      / [get]
func Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}

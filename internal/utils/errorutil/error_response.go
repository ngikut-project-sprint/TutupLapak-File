package errorutil

import "github.com/labstack/echo/v4"

func SendErrorResponse(ctx echo.Context, msg string, statusCode int) error {
	return ctx.JSON(statusCode, map[string]string{
		"error": msg,
	})
}

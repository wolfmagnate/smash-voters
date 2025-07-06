package handler

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func parseID(c echo.Context, paramName string) (int, error) {
	idStr := c.Param(paramName)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}
	return id, nil
}

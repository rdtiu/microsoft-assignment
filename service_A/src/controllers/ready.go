package controllers

import (
    "net/http"
    "github.com/labstack/echo/v4"
)


func IsReady(c echo.Context) (err error) {
    return c.JSON(http.StatusOK, "System is Ready")
}

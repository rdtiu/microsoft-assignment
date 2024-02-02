package controllers

import (
    "net/http"
    "github.com/labstack/echo/v4"
)


func IsLive(c echo.Context) (err error) {
    return c.JSON(http.StatusOK, "System is Live")
}

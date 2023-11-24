package main

import (
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
)

func errorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := "Internal Server Error, check server logs for more information"

	if err == echo.ErrMethodNotAllowed {
		code = http.StatusMethodNotAllowed
		message = "Method not allowed"
	}

	req := c.Request()

	he, ok := err.(*echo.HTTPError)
	if ok {
		code = he.Code
		message = fmt.Sprintf("%s", he.Message)
	}

	log.Warnf("[%s - %s] %s", req.Method, req.URL.String(), err)

	c.String(code, message)
}

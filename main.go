package main

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
)

var (
	secureToken        string
	port               string = "3000"
	webhooks                  = make(map[string]string)
	templatesDirectory string = "templates"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	e.POST("/:name", func(c echo.Context) error {

		if c.QueryParam("token") != secureToken {
			c.String(http.StatusUnauthorized, "Unauthorized")
		}

		name := c.Param("name")
		if name == "" {
			log.Error("Missing name")
			return c.String(http.StatusBadRequest, "Missing name")
		}

		webhookURL, ok := webhooks[name]
		if !ok {
			log.Error("Unknown name", "name", name)
			return c.String(http.StatusBadRequest, "Unknown name")
		}

		var body RequestBody
		c.Bind(&body)

		log.Debug("Request received!", "url", webhookURL, "body", body)

		return c.String(http.StatusOK, "Request received!")
	})

	e.Start(":" + port)
}

type RequestBody struct {
	ID             int     `json:"id"`
	ActivityTime   float64 `json:"activityTime"`
	ActivityType   string  `json:"activityType"`
	StatusCode     string  `json:"statusCode"`
	Status         string  `json:"status"`
	ActivityResult string  `json:"activityResult"`
	UserID         int     `json:"userId"`
	Message        string  `json:"message"`
	Type           string  `json:"type"`
	Data           `json:"data"`
}

type Data struct {
	Message struct {
		Code   string `json:"code"`
		Params struct {
			ClientID     string `json:"clientId"`
			ClientName   string `json:"clientName"`
			AppUserName  string `json:"appUserName"`
			AppUserID    string `json:"appUserId"`
			AppUserEmail string `json:"appUserEmail"`
		} `json:"params"`
	} `json:"message"`
}

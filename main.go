package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"text/template"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
)

var (
	secureToken        string
	logLevel           log.Level = log.InfoLevel
	port               string    = "3000"
	webhooks                     = make(map[string]string)
	templatesDirectory string    = "templates"
	tmpl               *template.Template
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.HTTPErrorHandler = errorHandler

	tmpl = template.New("").Funcs(template.FuncMap{
		"expand": func(p any) string {
			return fmt.Sprintf("%+v", p)
		},
	})
	tmpl = template.Must(tmpl.ParseGlob(templatesDirectory + "/*.html"))

	e.POST("/:name", func(c echo.Context) error {

		// Check token
		if c.QueryParam("token") != secureToken {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}

		name := c.Param("name")
		if name == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Missing name")
		}

		// Matching webhook with name
		webhookURL, ok := webhooks[name]
		if !ok {
			return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Webhook not found for '%s'", name))
		}

		// Parsing body from NinjaOne RMM
		var body RequestBody
		if logLevel == log.DebugLevel {
			// In debug mode, we want to see the raw body for developpement purpose
			b, _ := io.ReadAll(c.Request().Body)
			log.Debugf("Webhook received for '%s' and payload:\n%s", name, b)
			err := json.Unmarshal(b, &body)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "Failed decoding body")
			}
		} else {
			// In production, we want to simply decode the body
			c.Bind(&body)
		}

		go sendTeams(webhookURL, body)

		return c.String(http.StatusOK, "OK")
	})

	log.Fatal(e.Start(":" + port))
}

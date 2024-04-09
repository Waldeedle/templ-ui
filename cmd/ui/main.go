package main

import (
	"encoding/json"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo"
	"github.com/waldeedle/templ-ui/ui"
)

func main() {
	e := echo.New()
	e.Debug = true

	e.Static("/static", "../../assets")

	e.GET("/", func(c echo.Context) error { return HTML(c, ui.Index()) })
	e.GET("/toast", func(c echo.Context) error {
		// Construct your toast message as a map
		toastMessage := map[string]map[string]string{
			"toast-show": {
				"type":        "default",
				"message":     "test",
				"description": "dest",
				"position":    "top-center",
			},
		}

		// Marshal the toastMessage into JSON
		toastMessageJson, err := json.Marshal(toastMessage)
		if err != nil {
			return err
		}

		// Set the HX-Trigger header with the marshaled JSON
		c.Response().Header().Set("HX-Trigger", string(toastMessageJson))

		// Return the response with the message "Hello, World!"
		return c.NoContent(http.StatusOK)
	})

	e.Logger.Fatal(e.Start(":8100"))
}

func HTML(c echo.Context, comp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return comp.Render(c.Request().Context(), c.Response().Writer)
}

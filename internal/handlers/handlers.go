package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pablovgdev/relink/internal/redirect"
)

func Redirect(c echo.Context) error {
	path := c.Request().URL.Path

	url, err := redirect.GetRedirectUrlByPath(path)
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusNotFound, "Not found")
	}

	return c.Redirect(http.StatusSeeOther, url)
}

type RedirectBody struct {
	Path string `json:"path"`
	URL  string `json:"url"`
}

func PostRedirect(c echo.Context) error {
	body := RedirectBody{}
	err := c.Bind(&body)
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "Bad request")
	}

	if body.Path == "" || body.URL == "" {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	if body.Path[0] != '/' {
		body.Path = "/" + body.Path
	}

	if body.URL[:4] != "http" {
		body.URL = "http://" + body.URL
	}

	err = redirect.PostRedirect(body.Path, body.URL)
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, "Internal server error")
	}

	return c.String(http.StatusCreated, "OK")
}

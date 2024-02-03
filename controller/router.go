package controller

import "github.com/labstack/echo/v4"

func NewEchoServer() *echo.Echo {
	e := echo.New()
	controllers(e)
	return e
}

func controllers(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Server is up")
	})

}

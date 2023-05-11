package http

import (
	"TodoAPI/docs"
	"TodoAPI/internal/factory"
	"fmt"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
	"os"
)

func Init(e *echo.Echo, f *factory.Factory) {
	docs.SwaggerInfo.Host = os.Getenv("BASE_URL")
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/", func(c echo.Context) error {
		message := fmt.Sprintf("TODO LIST API")
		return c.String(http.StatusOK, message)
	})
}

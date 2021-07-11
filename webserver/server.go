package webserver

import (
	"net/http"

	"github.com/brunovivaldodev/strat-with-go/data"
	"github.com/labstack/echo/v4"
)

type WebServer struct {
}

func NewWebServer() *WebServer {
	return &WebServer{}
}
func (w WebServer) Serve() {
	e := echo.New()
	e.GET("/product", w.getAll)
	e.Logger.Fatal(e.Start(":8585"))
}

func (w WebServer) getAll(c echo.Context) error {
	return c.JSON(http.StatusOK, data.Products)
}

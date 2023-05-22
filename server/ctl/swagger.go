package ctl

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

// SwaggerController is the HTTP controller for microservice
// realtime documentation using swagger docs
type SwaggerController struct {
	baseController
	user     string
	password string
}

// NewSwaggerController creates a new Swagger route controller with given username and passord
func NewSwaggerController(user, password string) SwaggerController {
	return SwaggerController{
		user:     user,
		password: password,
	}
}

// compilation time interface implementation check
var _ HTTPRegistrar = (*SwaggerController)(nil)

// redirectToIndex redirect user browser to swagger index.html document
func (s SwaggerController) redirectToIndex(c echo.Context) error {
	dst := c.Request().RequestURI + "/index.html"
	return c.Redirect(http.StatusTemporaryRedirect, dst)
}

// Register adds new endpoints to current server handler
func (s SwaggerController) Register(e *echo.Group) {
	e.GET("/docs/*", echoSwagger.WrapHandler)
	e.GET("/docs", s.redirectToIndex)
	e.GET("/docs/", s.redirectToIndex)
}

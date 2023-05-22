package ctl

import "github.com/labstack/echo/v4"

// baseController is the basic datatype used to store all common
// properties among all HTTP controllers
type baseController struct{}

// HTTPRegistrar is the interface that enables to any implementer to register
// http endpoints in the given server
type HTTPRegistrar interface {
	Register(e *echo.Group)
}

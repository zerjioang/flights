package ctl

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"time"
)

var (
	host, _ = os.Hostname()
)

// HealthController is the HTTP controller for microservice
// health check related data
type HealthController struct {
	baseController
}

// compilation time interface implementation check
var _ HTTPRegistrar = (*HealthController)(nil)

// NewHealthController is a constructor like function that returns
// a new HealthController
func NewHealthController() HealthController {
	return HealthController{}
}

// healthCheck makes a simple server health check
//
//	@Summary		Get service API status
//	@Description	get the Flights service API status data
//	@Tags			healthCheck
//	@ID				healthCheck-get
//	@Produce		json
//	@Success		200	{object} map[string]interface{}
//	@Router			/v1/health [get]
func (ctl HealthController) healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":     "pong",
		"hostname": host,
		"ts":       time.Now().Unix(),
	})
}

// Register adds new endpoints to current server handler
func (ctl HealthController) Register(e *echo.Group) {
	e.GET("/health", ctl.healthCheck)
}

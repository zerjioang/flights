package ctl

import "github.com/labstack/echo/v4"

// FlightController is the HTTP controller for flight related data
type FlightController struct {
	baseController
}

// compilation time interface implementation check
var _ HTTPRegistrar = (*FlightController)(nil)

// NewFlightController is a constructor like function that returns
// a new FlightController
func NewFlightController() FlightController {
	return FlightController{}
}

// calculate runs a flight track calculation
// @Summary run a flight track calculation of given passenger
// @Description run a flight track calculation of given passenger
// @Tags FlightsCalculate
// @ID flightsCalculate-GET
// @Accept json
// @Produce json
// @Param   flightData	body [][]string true "Flight data"
// @Success 200 {object} []string
// @Failure 500 {object} map[string]interface{}	"Internal Server Error"
// @Router /v1/calculate [post]
func (ctl FlightController) calculate(ctx echo.Context) error {
	return nil
}

// Register adds new endpoints to current server handler
func (ctl FlightController) Register(e *echo.Group) {
	e.POST("/calculate", ctl.calculate)
}

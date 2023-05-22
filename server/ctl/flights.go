// Copyright zerjioang. 2023 All Rights Reserved.
// Licensed under the MIT
// you may not use this file except in compliance with the License.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package ctl

import (
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zerjioang/flights/core"
	"github.com/zerjioang/flights/server/datatype"
)

var (
	errInvalidJson = errors.New("invalid JSON request content")
)

// FlightController is the HTTP controller for flight related data
type FlightController struct {
	_      baseController
	solver core.FlightSolver
}

// compilation time interface implementation check
var _ HTTPRegistrar = (*FlightController)(nil)

// NewFlightController is a constructor like function that returns
// a new FlightController
func NewFlightController(solver core.FlightSolver) FlightController {
	if solver == nil {
		panic("solver not defined. Please check your calls to NewFlightController()")
	}
	return FlightController{solver: solver}
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
	var data datatype.FlightData
	// read incoming http data
	if err := data.Load(ctx.Request().Body); err != nil {
		log.Println("failed to unmarshal json body due to:", err)
		return errInvalidJson
	}
	passengerFlight, err := ctl.solver.Solve(&data)
	if err != nil {
		log.Println("failed to solve flight tracking data due to:", err)
		return err
	}
	return ctx.JSON(http.StatusOK, passengerFlight)
}

// Register adds new endpoints to current server handler
func (ctl FlightController) Register(e *echo.Group) {
	e.POST("/calculate", ctl.calculate)
}

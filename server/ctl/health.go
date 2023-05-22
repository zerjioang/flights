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
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

var (
	host, _ = os.Hostname()
)

// HealthController is the HTTP controller for microservice
// health check related data
type HealthController struct {
	_ baseController
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

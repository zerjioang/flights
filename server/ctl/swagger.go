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

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// SwaggerController is the HTTP controller for microservice
// realtime documentation using swagger docs
type SwaggerController struct {
	_        baseController
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

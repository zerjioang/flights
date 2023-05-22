//
// Copyright zerjioang. 2023 All Rights Reserved.
// Licensed under the MIT
// you may not use this file except in compliance with the License.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package core

import "github.com/zerjioang/flights/server/datatype"

// FlightSolver is the interface definition of the flight tracking solver implementation
// Optimizing the solver immplementation and swapping in the interface usage will affect
// to entire application
type FlightSolver interface {
	Solve(data *datatype.FlightData) (*datatype.Flight, error)
}

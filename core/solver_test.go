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

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/easypprof"
	"github.com/zerjioang/flights/server/datatype"
)

func TestBasicSolver_Solve(t *testing.T) {
	t.Run("solver", func(t *testing.T) {
		type testdata struct {
			Name string
			In   func() *datatype.FlightData
			Out  *datatype.Flight
			Err  error
		}
		testcases := []testdata{
			{
				Name: "size-1",
				In: func() *datatype.FlightData {
					var data datatype.FlightData
					_ = data.Load(strings.NewReader(`[["SFO", "EWR"]]`))
					return &data
				},
				Out: datatype.NewFlight("SFO", "EWR"),
				Err: nil,
			},
			{
				Name: "size-2",
				In: func() *datatype.FlightData {
					var data datatype.FlightData
					_ = data.Load(strings.NewReader(`[["ATL", "EWR"], ["SFO", "ATL"]]`))
					return &data
				},
				Out: datatype.NewFlight("SFO", "EWR"),
				Err: nil,
			},
			{
				Name: "size-4",
				In: func() *datatype.FlightData {
					var data datatype.FlightData
					_ = data.Load(strings.NewReader(`[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]`))
					return &data
				},
				Out: datatype.NewFlight("SFO", "EWR"),
				Err: nil,
			},
		}
		for _, testcase := range testcases {
			t.Run(testcase.Name, func(t *testing.T) {
				s := NewBasicSolver()
				solution, err := s.Solve(testcase.In())
				assert.NoError(t, err)
				assert.NotNil(t, solution)
				t.Log("solution is:", solution)
				assert.Equal(t, solution, testcase.Out)
			})
		}

	})
	t.Run("test-unmarshaller", func(t *testing.T) {
		s := NewBasicSolver()
		var data datatype.FlightData
		assert.NoError(t, data.Load(strings.NewReader(`[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]`)))
		solution, err := s.Solve(&data)
		assert.NoError(t, err)
		assert.NotNil(t, solution)
		t.Log("solution is:", solution)
		assert.Equal(t, solution, datatype.NewFlight("SFO", "EWR"))
	})
}

func TestProfiler(t *testing.T) {
	t.Run("profile-multiple", func(t *testing.T) {
		easypprof.Profile(t, 500000, func() {
			s := NewBasicSolver()
			var data datatype.FlightData
			assert.NoError(t, data.Load(strings.NewReader(`[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]`)))
			solution, err := s.Solve(&data)
			assert.NoError(t, err)
			assert.NotNil(t, solution)
		})
	})
}

func BenchmarkBasicSolver_Solve(b *testing.B) {
	// baseline
	// BenchmarkBasicSolver_Solve/sample-8         	 1876402	       629.4 ns/op	   1.59 MB/s	     160 B/op	      11 allocs/op
	// converting city names from Name datatype to string
	// BenchmarkBasicSolver_Solve/sample-8         	 3122499	       368.5 ns/op	   2.71 MB/s	     384 B/op	       6 allocs/op
	// after optimizing the algorithm. first iteration
	// BenchmarkBasicSolver_Solve/sample-8         	 3717075	       319.8 ns/op	   3.13 MB/s	     320 B/op	       3 allocs/op
	// after adding dynamic destination resolver
	// BenchmarkBasicSolver_Solve/sample-8         	 3634096	       326.0 ns/op	   3.07 MB/s	      64 B/op	       1 allocs/op
	b.Run("sample", func(b *testing.B) {
		s := NewBasicSolver()
		var data datatype.FlightData
		assert.NoError(b, data.Load(strings.NewReader(`[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]`)))
		easypprof.Bench(b, func() {
			solution, err := s.Solve(&data)
			assert.NoError(b, err)
			assert.NotNil(b, solution)
		})
	})
}

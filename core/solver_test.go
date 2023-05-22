package core

import (
	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/easypprof"
	"github.com/zerjioang/flights/server/datatype"
	"testing"
)

func TestBasicSolver_Solve(t *testing.T) {
	t.Run("solver", func(t *testing.T) {
		type testdata struct {
			Name string
			In   datatype.FlightData
			Out  *datatype.Flight
			Err  error
		}
		testcases := []testdata{
			{
				Name: "size-1",
				In: datatype.FlightData{
					datatype.NewFlight("SFO", "EWR"),
				},
				Out: datatype.NewFlight("SFO", "EWR"),
				Err: nil,
			},
			{
				Name: "size-2",
				In: datatype.FlightData{
					datatype.NewFlight("ATL", "EWR"),
					datatype.NewFlight("SFO", "ATL"),
				},
				Out: datatype.NewFlight("SFO", "EWR"),
				Err: nil,
			},
			{
				Name: "size-4",
				In: datatype.FlightData{
					datatype.NewFlight("ATL", "EWR"),
					datatype.NewFlight("SFO", "ATL"),
					datatype.NewFlight("GSO", "IND"),
					datatype.NewFlight("ATL", "GSO"),
				},
				Out: datatype.NewFlight("SFO", "EWR"),
				Err: nil,
			},
		}
		for _, testcase := range testcases {
			t.Run(testcase.Name, func(t *testing.T) {
				s := NewBasicSolver()
				solution, err := s.Solve(testcase.In)
				assert.NoError(t, err)
				assert.NotNil(t, solution)
				t.Log("solution is:", solution)
				assert.Equal(t, solution, testcase.Out)
			})
		}

	})
	t.Run("profile-multiple", func(t *testing.T) {
		easypprof.Profile(t, 5000000, func() {
			s := NewBasicSolver()
			solution, err := s.Solve(datatype.FlightData{
				datatype.NewFlight("ATL", "EWR"),
				datatype.NewFlight("SFO", "ATL"),
				datatype.NewFlight("GSO", "IND"),
				datatype.NewFlight("ATL", "GSO"),
			})
			assert.NoError(t, err)
			assert.NotNil(t, solution)
		})
	})
}

func BenchmarkBasicSolver_Solve(b *testing.B) {
	b.Run("sample", func(b *testing.B) {

	})
}

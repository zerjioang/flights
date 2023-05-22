package core

import "github.com/zerjioang/flights/server/datatype"

type BasicSolver struct {
}

var (
	// compilation time interface implementation check
	_ FlightSolver = (*BasicSolver)(nil)
)

func NewBasicSolver() FlightSolver {
	return &BasicSolver{}
}

func (b BasicSolver) Solve(data datatype.FlightData) (*datatype.Flight, error) {
	return datatype.NewFlight("SFO", "EWR"), nil
}

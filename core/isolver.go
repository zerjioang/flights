package core

import "github.com/zerjioang/flights/server/datatype"

// FlightSolver is the interface definition of the flight tracking solver implementation
// Optimizing the solver immplementation and swapping in the interface usage will affect
// to entire application
type FlightSolver interface {
	Solve(data datatype.FlightData) (*datatype.Flight, error)
}

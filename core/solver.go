package core

import (
	"errors"
	"fmt"
	"log"

	"github.com/zerjioang/flights/server/datatype"
)

var (
	// errNoData is the error returned when no input data is provided
	errNoData = errors.New("no flight data provided")
)

type BasicSolver struct {
}

var (
	// compilation time interface implementation check
	_ FlightSolver = (*BasicSolver)(nil)
)

// NewBasicSolver is a constructor like function that returns a new FlightSolver as BasicSolver
func NewBasicSolver() FlightSolver {
	return &BasicSolver{}
}

// Solve executes current flight tracking solver algorithm
func (b BasicSolver) Solve(data *datatype.FlightData) (*datatype.Flight, error) {
	verbose := false
	if data == nil || len(data.Input) == 0 {
		return nil, errNoData
	}
	if len(data.Input) == 0 {
		return data.Input[0], nil
	}
	if len(data.Input) == 1 {
		return data.Input[0], nil
	}
	var possibleOrigins = map[string]*datatype.Path{}
	var possibleDestinations = map[string]*datatype.Path{}

	for i := 0; i < len(data.Input); i++ {
		flight := data.Input[i]
		// store current city in analyzed points if not previously created
		fromId := flight.From.String()
		departureCity := data.Cities[fromId]
		// update status
		departureCity.To = data.Cities[flight.To.String()]
		dstId := flight.To.String()
		destinationCity := data.Cities[dstId]
		// update status
		destinationCity.From = data.Cities[flight.From.String()]
		// add current city as possible flight origin
		if departureCity.From == nil {
			possibleOrigins[fromId] = departureCity
		}
		delete(possibleOrigins, destinationCity.CityName)
		// add current city as possible flight destination
		if destinationCity.To == nil {
			possibleDestinations[dstId] = destinationCity
		}
		delete(possibleDestinations, departureCity.CityName)
	}

	if len(possibleOrigins) > 1 {
		log.Fatal("warning: multiple possible origins detected")
	}
	if len(possibleDestinations) > 1 {
		log.Fatal("warning: multiple possible destinations detected")
	}

	var src *datatype.Path
	for _, flight := range possibleOrigins {
		src = flight
		break
	}
	var dst *datatype.Path
	for _, flight := range possibleDestinations {
		dst = flight
		break
	}
	if verbose {
		fmt.Println("Origin:", src.CityName)
		fmt.Println("Destination:", dst.CityName)
	}
	return datatype.NewFlight(src.CityName, dst.CityName), nil
}

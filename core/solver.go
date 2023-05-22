package core

import (
	"errors"
	"fmt"
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
func (b BasicSolver) Solve(data datatype.FlightData) (*datatype.Flight, error) {
	if data == nil {
		return nil, errNoData
	}
	if len(data) == 1 {
		return data[0], nil
	}
	var analyzedPoints = map[string]*datatype.Flight{}
	var possibleOrigins = map[string]*datatype.Flight{}
	for i := 0; i < len(data); i++ {
		point := data[i]
		// store current city in analyzed points if not previously created
		if _, ok := analyzedPoints[point.From.String()]; !ok {
			fromId := point.From.String()
			prevPoint := datatype.NewFlight("???", fromId)
			point.PrevHop = prevPoint
			// update status
			prevPoint.PrevHop = lookupPoint(analyzedPoints, point.From.String(), nil)
			prevPoint.NextHop = lookupPoint(analyzedPoints, point.To.String(), point)
			// store the point
			analyzedPoints[fromId] = point
			// add current point as possible flight origin
			possibleOrigins[fromId] = point
		}
		if p, ok := analyzedPoints[point.To.String()]; !ok {
			dstId := point.To.String()
			// now create virtual destination hop if origin/destination is not available 'yet'
			// example
			// current: [["ATL", "EWR"]
			// next hop: [EWR, UNKNOWN]
			nextPoint := datatype.NewFlight(dstId, "???")
			point.NextHop = nextPoint
			// update status
			nextPoint.PrevHop = lookupPoint(analyzedPoints, point.From.String(), point)
			nextPoint.NextHop = lookupPoint(analyzedPoints, point.To.String(), nil)
			analyzedPoints[dstId] = nextPoint
			// remove this point from possible origins (conditional exists)
			delete(possibleOrigins, dstId)
		} else {
			// update existing node hops
			if p.PrevHop != nil && p.PrevHop.From.String() == "???" {
				point.NextHop = p
				delete(possibleOrigins, p.From.String())
			}
		}
	}
	if len(possibleOrigins) > 1 {
		fmt.Println("warning: multiple possible origins detected")
	}
	var origin *datatype.Flight
	for _, flight := range possibleOrigins {
		origin = flight
		break
	}
	return resolveTrip(origin)
}

func lookupPoint(data map[string]*datatype.Flight, id string, fallback *datatype.Flight) *datatype.Flight {
	p, ok := data[id]
	if ok {
		return p
	}
	return fallback
}

func resolveTrip(origin *datatype.Flight) (*datatype.Flight, error) {
	if origin == nil {
		return nil, errors.New("cannot detected trip origin")
	}
	//fmt.Println("Origin:", origin.From)
	var dst *datatype.Flight
	currentPoint := origin
	// to avoid infinite loops in case of bad inputs
	maxHopsAllowed := 100
	iter := 0
	for currentPoint != nil && iter < maxHopsAllowed {
		dst = currentPoint
		currentPoint = currentPoint.NextHop
		iter++
	}
	if dst == nil {
		return nil, errors.New("could not find a valid solution for given input")
	}
	//fmt.Println("Passenger flight connection numbers: ", iter)
	//fmt.Println("Destination:", dst.From)
	return datatype.NewFlight(origin.From.String(), dst.From.String()), nil
}

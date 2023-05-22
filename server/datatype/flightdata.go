package datatype

import (
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io"
)

// FlightData is the flight data model of all flights for a given passenger
// examples:
// [["SFO", "EWR"]]                                                 => ["SFO", "EWR"]
// [["ATL", "EWR"], ["SFO", "ATL"]]                                 => ["SFO", "EWR"]
// [["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]] => ["SFO", "EWR"]
type FlightData []*Flight

func (d *FlightData) Load(reader io.ReadCloser) error {
	var dst [][]string
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	if err := json.NewDecoder(reader).Decode(&dst); err != nil {
		return err
	}
	for _, item := range dst {
		if len(item) == 2 {
			f := NewFlight(item[0], item[1])
			*d = append(*d, f)
		} else {
			return errors.New("invalid flight data found")
		}
	}
	return nil
}

// Flight is the content of one single item that represents a passenger flight
type Flight struct {
	From  Name
	To    Name
	Valid bool
	// used while analyzing content
	PrevHop *Flight
	NextHop *Flight
}

// Name is the name of the flight
type Name [3]byte

// NewFlight creates a new flight and parses its input data
func NewFlight(from, to string) *Flight {
	fromB, okFrom := parseName(from)
	toB, okTo := parseName(to)
	return &Flight{
		From:  fromB,
		To:    toB,
		Valid: okFrom && okTo,
	}
}

func (f *Flight) UnmarshalJSON(data []byte) error {
	if len(data) > 0 {
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		var inputFormat [2]string
		err := json.Unmarshal(data, &inputFormat)
		if err != nil {
			return err
		}
		fromB, okFrom := parseName(inputFormat[0])
		toB, okTo := parseName(inputFormat[1])
		f.From = fromB
		f.To = toB
		f.Valid = okFrom && okTo
	}
	return nil
}

// MarshalJSON overrides default JSON format for Flight data model
func (f *Flight) MarshalJSON() ([]byte, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var responseFormat [2]string
	responseFormat[0] = f.From.String()
	responseFormat[1] = f.To.String()
	return json.Marshal(responseFormat)
}

func (f *Flight) Id() string {
	return fmt.Sprintf("%s->%s", f.From, f.To)
}

// parseName validates and parses input name
func parseName(name string) (Name, bool) {
	if len(name) != 3 {
		return Name{}, false
	}
	// safe bytes to string conversion
	return Name([]byte(name)), true
}

func (n Name) String() string {
	// safe string to byte conversion
	return string([]byte{n[0], n[1], n[2]})
}

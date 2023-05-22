package datatype

import (
	"sync"
)

var flightPool = sync.Pool{
	// New creates an object when the pool has nothing available to return.
	// New must return an interface{} to make it flexible. You have to cast
	// your type after getting it.
	New: func() interface{} {
		// Pools often contain things like *bytes.Buffer, which are
		// temporary and re-usable.
		return &Flight{}
	},
}

func AdquireFlight() *Flight {
	return flightPool.Get().(*Flight)
}

func ReleaseFlight(f *Flight) {
	f.Reset()
	flightPool.Put(f)
}

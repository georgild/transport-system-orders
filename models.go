package main

import "gopkg.in/mgo.v2/bson"

// Pair : represents a pair of elements
type Pair struct {
	elem1, elem2 int
}

// Seat : represents seat
type Seat struct {
	Row int
	Col int
}

// Order : represents order
type Order struct {
	_id     bson.ObjectId `bson:"_id,omitempty"`
	RouteID string
	User    string
	Seats   []Seat
}

// ErrorString : used for error handling
type ErrorString struct {
	msg string
}

func (e *ErrorString) Error() string {
	return e.msg
}

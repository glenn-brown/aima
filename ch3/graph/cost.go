package graph

import . "github.com/glenn-brown/builtin"

type Cost interface {
	Add(interface{}) interface{}
	Less(interface{}) bool
}

// NewCost() returns a new Cost interface for a go builtin scalar type.
func NewCost(i interface{}) Cost {
	return Augment(i).(Cost)
}

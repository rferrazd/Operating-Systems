package main

import "sync" // mutex from here
// no main function. compiler looks for the main function

// a chopstick is like a resource
type ChopStick struct {
	sync.Mutex
	// another type inside struct, inherent all properties from that type.
	// chopstick inherets all attributes
	// Mutually exclusive resource
	ID int // member variable
}

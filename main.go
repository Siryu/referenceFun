package main

import (
	"log"
)

type Bottle struct {
	CapColor     string
	BottleHeight int
}

func main() {
	differentCreations()
	dereferencePointer()
	passingPointer()
}

func differentCreations() {
	// use new to instantiate a new bottle with default values and get a pointer back
	bottle := new(Bottle)
	log.Printf("capcolor: %v", bottle.CapColor)
	log.Printf("height: %v", bottle.BottleHeight)

	// use & to get a pointer back
	bottle2 := &Bottle{}
	log.Printf("2 capcolor: %v", bottle2.CapColor)
	log.Printf("2 height: %v", bottle2.BottleHeight)
}

func dereferencePointer() {
	// pass a dereferenced pointer to get a copy, not the object
	bottle3 := &Bottle{
		CapColor:     "red",
		BottleHeight: 9,
	}
	testDereference(*bottle3)
	// will be the same as when we instantiated since we dereference before we pass in
	log.Printf("3 capcolor: %v", bottle3.CapColor)
	log.Printf("3 height: %v", bottle3.BottleHeight)
}

func passingPointer() {
	// pass a pointer to change the original object
	bottle4 := &Bottle{
		CapColor:     "yellow",
		BottleHeight: 15,
	}
	testPassPointer(bottle4)
	// will be changed to blue and 11 because the reference data was changed
	log.Printf("4 capcolor: %v", bottle4.CapColor)
	log.Printf("4 height: %v", bottle4.BottleHeight)
}

func testPassPointer(b *Bottle) {
	b.CapColor = "blue"
	b.BottleHeight = 11
}

func testDereference(b Bottle) {
	// adjust the values of the bottle that is passed in
	b.CapColor = "green"
	b.BottleHeight = 10
}

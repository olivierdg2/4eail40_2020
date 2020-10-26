<<<<<<< HEAD
// Package coord will contain types and logic for handling chess biard coordinates.
=======
// Package coord will contain types and logic for handling chess board coordinates.
>>>>>>> prof/master
package coord

import "fmt"

<<<<<<< HEAD
//ChessCoordinates is an interface for an abstract Coordinates.
type ChessCoordinates interface {
	fmt.Stringer // "A7"
	//Coord get n'th coordinates comp.
	// Start with 0th coordinate.
	Coord(n int) (int, error)

	// Vector
	//
=======
// ChessCoordinates is an interface for an abstract Coordinates.
type ChessCoordinates interface {
	fmt.Stringer // "A7"
	// Coord gets n'th coordinate comp.
	// Start with 0th coordinate.
	// Returns an error if coordinate is inexistant.
	Coord(n int) (int, error)

	// Vector
	// Add()
	// Multiply()
>>>>>>> prof/master
}

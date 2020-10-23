// Package coord will contain types and logic for handling chess biard coordinates.
package coord

import "fmt"

//ChessCoordinates is an interface for an abstract Coordinates. 
type ChessCoordinates interface {
	fmt.Stringer // "A7"
	//Coord get n'th coordinates comp.
	// Start with 0th coordinate.
	Coord(n int) (int, error)
}
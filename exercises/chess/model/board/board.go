// Package board will contain types and logic for handling chess board(s).
package board

import (
	"fmt"
	"github.com/olivierdg2/4eail40_2020/exercises/chess/model/piece"
	"github.com/olivierdg2/4eail40_2020/exercises/chess/model/coord"
)

//TODO Implement type Board

//Board is an interface to a chess board,
// It defines methods for handling pieces on it.
type Baord interface {
	fmt.Stringer
	// PieceAt retrives piece at given coordinates.
	// Returns nil if no piece was found.
	PieceAt(at coord.ChessCoordinates) piece.Piece
	// MovePiece moves a piece from given coordinates to
	// given coordinates.
	MovePiece(from, to coord.ChessCoordinates) error
	// PlacePieceAt places a given piece at given location.
	// Returns an error if destination was occuplied.
	PlacePieceAt(p piece.Piece, at coord.ChessCoordinates) error
}

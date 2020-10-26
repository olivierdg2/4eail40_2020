<<<<<<< HEAD
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
=======
package board

import (
	"github.com/jglouis/4eail40_2020/exercises/chess/model/coord"
	"github.com/jglouis/4eail40_2020/exercises/chess/model/piece"
)

// Classic 8x8 Chess board
type Classic [8][8]piece.Piece

func (c *Classic) String() string {
	panic("not implemented") // TODO: Implement
}

// PieceAt retrievves piece at give coordinates.
// Returns nil if no piece was found.
func (c *Classic) PieceAt(at coord.ChessCoordinates) piece.Piece {
	panic("not implemented") // TODO: Implement
}

// MovePiece moves a piece from given coordinates to
// given coordinates.
// Returns an error if destination was occupied.
func (c *Classic) MovePiece(from coord.ChessCoordinates, to coord.ChessCoordinates) error {
	panic("not implemented") // TODO: Implement
}

// PlacePieceAt places a given piece at given location.
// Returns an error if destination was occupied.
func (c *Classic) PlacePieceAt(p piece.Piece, at coord.ChessCoordinates) error {
	panic("not implemented") // TODO: Implement
>>>>>>> prof/master
}

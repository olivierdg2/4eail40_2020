<<<<<<< HEAD
=======
// Package piece will handle game pieces.
>>>>>>> prof/master
package piece

import (
	"fmt"

<<<<<<< HEAD
	"github.com/olivierdg2/4eail40_2020/exercises/chess/model/coord"
	"github.com/olivierdg2/4eail40_2020/exercises/chess/model/player"
=======
	"github.com/jglouis/4eail40_2020/exercises/chess/model/coord"
	"github.com/jglouis/4eail40_2020/exercises/chess/model/player"
>>>>>>> prof/master
)

// Piece represents a game piece
type Piece interface {
	fmt.Stringer
<<<<<<< HEAD
	// Color returns
	Color() player.Color
	// Moves returns
	Moves() map[coord.ChessCoordinates]bool
=======
	// Color returns the appartenance.
	Color() player.Color
	// Moves returns a set of valid move.
	Moves(isCapture bool) map[coord.ChessCoordinates]bool
>>>>>>> prof/master
}

package piece

import (
	"fmt"

	"github.com/olivierdg2/4eail40_2020/exercises/chess/model/coord"
	"github.com/olivierdg2/4eail40_2020/exercises/chess/model/player"
)

// Piece represents a game piece
type Piece interface {
	fmt.Stringer
	// Color returns
	Color() player.Color
	// Moves returns
	Moves() map[coord.ChessCoordinates]bool
}

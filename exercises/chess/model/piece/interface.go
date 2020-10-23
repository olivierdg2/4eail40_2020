package piece

import (
	"fmt"

	"github.com/olivierdg2/4eail40_2020/exercices/chess/model/coord"
	"github.com/olivierdg2/4eail40_2020/exercices/chess/model/player"
)

// Piece represents a game piece
type Piece interface {
	fmt.Stringer
	// Color returns
	Color() player.Color
	// Moves returns
	Moves() []coord.ChessCoordinates
}

package coord

import "fmt"

// Cartesian represents a set of cartesian coordinates, x and y.
type Cartesian struct {
	x, y int
}

// NewCartesian is the constructor for Cartesian
func NewCartesian(x, y int) Cartesian {
	return Cartesian{x, y}
}

// Coord x if n == 0, y if n == 1
func (c Cartesian) Coord(n int) (int, error) {
	switch n {
	case 0:
		return c.x, nil
	case 1:
		return c.y, nil
	default:
		return 0, fmt.Errorf("unknown coord component %d", n)
	}
}

func (c Cartesian) String() string {
	panic("TODO")
}

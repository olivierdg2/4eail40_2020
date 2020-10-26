package coord

import "fmt"

<<<<<<< HEAD
// Cartesian represents a set of cartesian coordinates, x and y.
=======
// Cartesian represent a set of cartesian coordinates, x and y.
>>>>>>> prof/master
type Cartesian struct {
	x, y int
}

<<<<<<< HEAD
// NewCartesian is the constructor for Cartesian
=======
// NewCartesian is the constructer for Cartesian.
>>>>>>> prof/master
func NewCartesian(x, y int) Cartesian {
	return Cartesian{x, y}
}

<<<<<<< HEAD
// Coord x if n == 0, y if n == 1
=======
// Coord returns x if n==0, y if n==1
>>>>>>> prof/master
func (c Cartesian) Coord(n int) (int, error) {
	switch n {
	case 0:
		return c.x, nil
	case 1:
		return c.y, nil
<<<<<<< HEAD
	default:
		return 0, fmt.Errorf("unknown coord component %d", n)
	}
}

func (c Cartesian) String() string {
	panic("TODO")
=======

	}
	return 0, fmt.Errorf("unknown coord component %d", n)
}

func (c Cartesian) String() string {
	return fmt.Sprintf("%c%d", 65+c.y, c.x+1)
>>>>>>> prof/master
}

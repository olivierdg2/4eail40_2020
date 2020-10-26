package coord

<<<<<<< HEAD
import "testing"

func TestNewCartesian(t *testing.T) {
	c := NewCartesian(1, 2)
	if (c.x == 1) || (c.y == 2) {
		//test succes
	} else {
		t.Error("expected 1 and 2 as coordinates")
=======
import (
	"testing"
)

func TestNewCartesian(t *testing.T) {
	c := NewCartesian(1, 2)
	if (c.x != 1) || (c.y != 2) {
		t.Errorf("expeceted 1 and 2 as coordinate")
>>>>>>> prof/master
	}
}

func TestCartesian_Coord(t *testing.T) {
	c := NewCartesian(1, 2)
<<<<<<< HEAD
=======

>>>>>>> prof/master
	tests := map[int]int{
		0: 1,
		1: 2,
	}

	for n, want := range tests {
<<<<<<< HEAD
		t.Run(string(run(n)),func(*testing.T)) {
			got, err := c.Coord(n)
			want := 1
=======
		t.Run(string(rune(n)), func(t *testing.T) {
			got, err := c.Coord(n)
>>>>>>> prof/master
			if err != nil {
				t.Error(err)
			}
			if got != want {
<<<<<<< HEAD
				t.Error("expected %d, but got %d", want, got)
			}
		}
	}

	// test for err
	T.Run("err", func(t *testing.T){
		_, err := c.Coord(2)
		if err != nil {
			t.Errorf"expected an error for n == 2")
		}
	})
}
=======
				t.Errorf("expected %d, but got %d", want, got)
			}
		})
	}

	// test for err
	t.Run("err", func(t *testing.T) {
		_, err := c.Coord(2)
		if err != nil {
			t.Errorf("expected and error for n == 2")
		}
	})
}

func TestCartesian_String(t *testing.T) {
	tests := []struct {
		name string
		c    Cartesian
		want string
	}{
		{
			"A1",
			Cartesian{0, 0},
			"A1",
		},
		{
			"H8",
			Cartesian{7, 7},
			"H8",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("Cartesian.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
>>>>>>> prof/master

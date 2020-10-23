package coord

import "testing"

func TestNewCartesian(t *testing.T) {
	c := NewCartesian(1, 2)
	if (c.x == 1) || (c.y == 2) {
		//test succes
	} else {
		t.Error("expected 1 and 2 as coordinates")
	}
}

func TestCartesian_Coord(t *testing.T) {
	c := NewCartesian(1, 2)
	tests := map[int]int{
		0: 1,
		1: 2,
	}

	for n, want := range tests {
		t.Run(string(run(n)),func(*testing.T)) {
			got, err := c.Coord(n)
			want := 1
			if err != nil {
				t.Error(err)
			}
			if got != want {
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

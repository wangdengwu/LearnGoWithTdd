//@Author: Dennis
//@Date: 2023/5/11

package shape

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("rectangle area", func(t *testing.T) {
		var shape Shape = Rectangle{
			Width:  8.0,
			Height: 9.0,
		}
		got := shape.Area()
		want := 72.0
		if want != got {
			t.Errorf("want %g got %g", want, got)
		}
	})

	t.Run("circle area", func(t *testing.T) {
		var shape Shape = Circle{
			Radius: 8.0,
		}
		got := shape.Area()
		want := math.Pi * 8 * 8
		if want != got {
			t.Errorf("want %g got %g", want, got)
		}
	})
}

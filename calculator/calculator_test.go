//@Author: Dennis
//@Date: 2023/5/11

package calculator

import (
	"fmt"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("test two int add", func(t *testing.T) {
		got := Add(1, 1)
		want := 2
		if want != got {
			t.Errorf("want %d got %d", want, got)
		}
	})
	t.Run("test add multi", func(t *testing.T) {
		got := Add(1, 2, 3)
		want := 6
		if want != got {
			t.Errorf("want %d got %d", want, got)
		}
	})
	t.Run("test two float add", func(t *testing.T) {
		got := Add(1.0, 1.2, 1.3)
		want := 3.5
		if want != got {
			t.Errorf("want %g got %g", want, got)
		}
	})
}

func TestDiv(t *testing.T) {
	t.Run("test two int div", func(t *testing.T) {
		got := Div(10, 2)
		want := 5.0
		if want != got {
			t.Errorf("want %g got %g", want, got)
		}
	})
	t.Run("test div 0", func(t *testing.T) {
		got := Div(1, 0)
		if got != math.Inf(10) {
			t.Errorf("not Inf got %g", got)
		}
	})
	t.Run("test 0/0", func(t *testing.T) {
		got := Div(0, 0)
		if got != got {
			fmt.Println(fmt.Sprintf("NaN != %v", got))
		} else {
			t.Errorf("not Inf got %g", got)
		}
	})
}

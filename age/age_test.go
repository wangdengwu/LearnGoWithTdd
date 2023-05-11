//@Author: Dennis
//@Date: 2023/5/11

package age

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestAge(t *testing.T) {
	t.Run("use slice", func(t *testing.T) {
		peoples := []People{{
			Name: "wo",
			Age:  1,
		}, {
			Name: "ni",
			Age:  2,
		}, {
			Name: "ta",
			Age:  3,
		}}
		IncSlice(peoples)
		_, _ = fmt.Fprintf(os.Stdout, "after inc %v\n", peoples)
		want := []People{{
			Name: "wo",
			Age:  2,
		}, {
			Name: "ni",
			Age:  3,
		}, {
			Name: "ta",
			Age:  4,
		}}
		if !reflect.DeepEqual(want, peoples) {
			t.Errorf("want %v got %v", want, peoples)
		}
	})
	t.Run("use map", func(t *testing.T) {
		peoples := map[string]People{
			"wo": {
				Name: "wo",
				Age:  1,
			},
			"ni": {
				Name: "ni",
				Age:  2,
			},
			"ta": {
				Name: "ta",
				Age:  3,
			},
		}
		IncMap(peoples)
		for _, v := range peoples {
			_, _ = fmt.Fprintf(os.Stdout, "after inc %v\n", v)
		}
		want := peoples
		if !reflect.DeepEqual(want, peoples) {
			t.Errorf("want %v got %v", want, peoples)
		}
	})

}

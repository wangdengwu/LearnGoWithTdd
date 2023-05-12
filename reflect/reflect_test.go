//@Author: Dennis
//@Date: 2023/5/12

package reflect

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"unsafe"
)

type Object struct {
	Int    int
	str    string `weee:"go"`
	Array  []int
	Map    map[string]string
	Struct struct{}
}

func (o Object) Method(_ int, s string) (string, error) {
	return s, nil
}

func TestReflect(t *testing.T) {
	t.Run("reflect struct", func(t *testing.T) {
		obj := Object{
			Int:    1,
			str:    "hello",
			Array:  []int{1, 2},
			Map:    map[string]string{"1": "1", "2": "2"},
			Struct: struct{}{},
		}
		tag, _ := reflect.TypeOf(&obj).Elem().FieldByName("str")
		tagValue := tag.Tag.Get("weee")
		_, _ = fmt.Fprintf(os.Stdout, "str tag is %s\n", tagValue)
		valueOf := reflect.ValueOf(&obj)
		i := valueOf.Elem().FieldByName("Int")
		i.SetInt(2)
		_, _ = fmt.Fprintf(os.Stdout, "int type is %s value is %d\n", i.Type().String(), i.Int())
		s := valueOf.Elem().FieldByName("str")
		reflect.NewAt(s.Type(), unsafe.Pointer(s.UnsafeAddr())).Elem().SetString("world")
		//s.SetString("world")
		_, _ = fmt.Fprintf(os.Stdout, "str type is %s value is %s\n", s.Type().String(), s.String())
		a := valueOf.Elem().FieldByName("Array")
		_, _ = fmt.Fprintf(os.Stdout, "array type is %s value is %v a[0] is %d\n", a.Type().String(), a, a.Index(0).Int())
		m := valueOf.Elem().FieldByName("Map")
		_, _ = fmt.Fprintf(os.Stdout, "map type is %s value is %v m['1'] is %s\n", m.Type().String(), m, m.MapIndex(reflect.ValueOf("1")).String())
		o := valueOf.Elem().FieldByName("Struct")
		_, _ = fmt.Fprintf(os.Stdout, "struct type is %s value is %v\n", o.Type().String(), o)
		md := valueOf.Elem().MethodByName("Method")
		result := md.Call([]reflect.Value{reflect.ValueOf(1), reflect.ValueOf("hello")})
		_, _ = fmt.Fprintf(os.Stdout, "method type is %s value is %v call result is %s\n", md.Type().String(), md, result[0])
	})
}

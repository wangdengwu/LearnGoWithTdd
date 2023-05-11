//@Author: Dennis
//@Date: 2023/5/11

package age

type People struct {
	Name string
	Age  int
}

func IncSlice(peoples []People) {
	for i := 0; i < len(peoples); i++ {
		peoples[i].Age++
	}
}

func IncMap(peoples map[string]People) {
	for k, v := range peoples {
		v.Age++
		peoples[k] = v
	}
}

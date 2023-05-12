//@Author: Dennis
//@Date: 2023/5/11

package calculator

type Number interface {
	int | float64
}

func Add[T Number](is ...T) (sum T) {
	for i := 0; i < len(is); i++ {
		sum += is[i]
	}
	return
}

func Div(i int, j int) float64 {
	return float64(i) / float64(j)
}

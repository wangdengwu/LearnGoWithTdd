//@Author: Dennis
//@Date: 2023/5/11

package helloworld

const defaultName = "world"

func Hello(name string) string {
	if name == "" {
		name = defaultName
	}
	return "hello " + name
}

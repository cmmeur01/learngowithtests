package itr

import "fmt"

func Repeater(str string, times int) string {
	repeated := ""
	for i := 0; i < times; i++ {
		repeated += str
	}

	return repeated
}

func main() {
	fmt.Println(Repeater("b", 7))
}

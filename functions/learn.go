package functions

import (
	"fmt"
	"strings"
)

func Learn() {
	fmt.Println(add(10, 5))
	fmt.Println(converter("hello", "geekgreen@gmail.com"))

	sum(10, 20, 30, 40, 50, 60, 70, 80, 90, 00)
}

func add(x, y int) int {
	return x + y
}

func sum(numbers ...int) {

	// fmt.Println(numbers)
	total := 0
	for _, v := range numbers {
		total += v
	}
	fmt.Println(total)
}

func converter(title, email string) (string, string) {
	s1 := strings.ToUpper(title)
	return s1, email
}

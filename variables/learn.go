package variables

import (
	"fmt"
	"strconv"
)

var fullname string = "Non"
var email string = "codenon@hotmail.com"
var c, python bool = true, false

const vat int = 7

func Learn() {
	fullname = "Nontawat Pomsanam"
	fmt.Println(email)
	fmt.Printf("Myname is %s\n", fullname)
	fmt.Println(c, python)

	companyName := "SuwanService998"
	isShow := true
	age := 10

	fmt.Println(companyName, isShow, age)
	fmt.Printf("%T\n", isShow)

	f := float64(age)
	fmt.Printf("%T\n", f)

	s := strconv.Itoa(vat)
	fmt.Println("int to String Vat is " + s)
}

package pointers

import "fmt"

func Learn() {
	x := 10
	fmt.Println(x)
	fmt.Println(&x)

	y := x
	fmt.Println(y)
	fmt.Println(&y)

	// var p *int = &x
	p := &x //pointer ชี้ไปที่ x reference
	fmt.Printf("pointer type is %T\n", p)
	fmt.Println(p)
	fmt.Println(*p) //dereference

	*p = 20
	fmt.Println(x)
	fmt.Println(*p)

	j := student{name: "Jame"}

	fmt.Println(j)

	setName(&j)
	fmt.Println(j)
}

type student struct {
	name string
}

func setName(std *student) {
	std.name = "Boy"
}

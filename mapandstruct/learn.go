package mapandstruct

import (
	"fmt"
	"math"
)

func Learn() {
	//Map (key , value)
	var m2 = make(map[string]int)
	m2["token"] = 10000
	m := map[string]int{"token": 10, "access": 20}

	m["token"] = 100

	fmt.Println(m)

	//loop map m
	for key, v := range m {
		fmt.Println(key, v)
		fmt.Printf("%v -> %v \n", key, v)
	}

	delete(m, "access")
	m["token2"] = 100
	fmt.Println(m)
	fmt.Println(m2)

	// struct
	type User struct {
		id       int
		username string
		password string
	}

	jhon := User{
		id:       1,
		username: "Jhony Depp",
		password: "1234",
	}

	jhon.password = "12345678901234567890"
	fmt.Println(jhon.username)

	users := []User{
		{
			id: 1, username: "bob", password: "1212",
		},
		{
			id: 1, username: "bob", password: "1212",
		},
	}

	fmt.Println(users[0].username)

	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

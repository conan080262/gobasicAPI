package mapandstruct

import "fmt"

func Learn() {
	//Map (key , value)

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

	// struct

}

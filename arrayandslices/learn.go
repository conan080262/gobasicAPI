package arrayandslices

import "fmt"

func Learn() {
	//Array
	var arr [2]string
	arr[0] = "A"
	arr[1] = "B"
	fmt.Println(arr)

	arr2 := [4]int{1, 2, 3}
	fmt.Println(arr2)

	//Slice
	var s []int = arr2[1:3]
	fmt.Println(s)

	s2 := arr2[1:3]
	s3 := []int{2, 3, 4, 5, 6, 7}
	fmt.Println(s2)
	fmt.Printf("len=%d cap=%d\n", len(s2), cap(s2))
	fmt.Println(s3)

	x := []int{1, 2}
	// y := []int{3, 4}
	a := []int{3, 4}
	z := append(x, a...)
	fmt.Println(z)
}

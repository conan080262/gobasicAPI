package flowcontrol

import (
	"fmt"
	"io/ioutil"
	"runtime"
)

func Learn() {
	score := 10
	if score == 10 {
		fmt.Println("good")
	} else {
		fmt.Println("try again")
	}

	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("MacOS")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Printf("%s\n", os)
	}

	b, err := ioutil.ReadFile("file.txt")

	if err == nil { //nil = success
		content := string(b)
		fmt.Println(content)
	} else {
		fmt.Println(err, b)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 5 {
			fmt.Println("BREAK IF")
			break //continue
		}
	}

}

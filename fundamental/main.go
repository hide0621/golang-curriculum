package main

import (
	"errors"
	"fmt"
)

func main() {

	i := 10

	if i == 10 {
		fmt.Println("i is 10")
	} else {
		fmt.Println("i is not 10")
	}

	if true {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	// 上と下のif文は同じ意味

	ok := true

	if ok {
		fmt.Println("true")
	}

	if !ok {
		fmt.Println("false")
	}

	if err := errors.New("error"); err != nil {
		fmt.Println(err)
	}

}

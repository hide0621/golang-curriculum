package main

import "fmt"

func if_switch_practice() {
	numberStr := "one"

	if numberStr == "one" {
		fmt.Println("Number is one")
	} else if numberStr == "two" {
		fmt.Println("Number is two")
	} else if numberStr == "three" {
		fmt.Println("Number is three")
	} else {
		fmt.Println("Unknown number")
	}

	// 上記のようなif文は、switch文で書き換えた方がスマート

	switch numberStr {
	case "one":
		fmt.Println("Number is one")
	case "two":
		fmt.Println("Number is two")
	case "three":
		fmt.Println("Number is three")
	default:
		fmt.Println("Unknown number")
	}

	// こんなswitch文も書ける

	name := "Taro"
	switch name {
	case "Taro", "Jiro":
		fmt.Println("I am Taro or Jiro")
	case "Saburo":
		fmt.Println("I am Saburo")
	default:
		fmt.Println("I am someone")
	}

	switch id := 2; id {
	case 1, 2:
		fmt.Println("id is 1 or 2")
	case 3:
		fmt.Println("id is 3")
	default:
		fmt.Println("id is unknown")
	}

	// switch文の条件式（case式）には、bool型の変数を指定できる
	switch x := 1; x > 0 {
	case true:
		fmt.Println("x is positive")
	case false:
		fmt.Println("x is negative")
	default:
		fmt.Println("x is zero")
	}
}

package main

import "fmt"

func main() {
	var num1 = 56
	var num2 = 4
	var result = num1 + num2
	fmt.Println(result)

	var floatNum = 56.5
	var result2 = floatNum + float64(num2)
	fmt.Println(result2)

	var stringType = "Gabriel"
	fmt.Println(stringType)
	fmt.Println(len(stringType))

	var myBoolean = false
	fmt.Println(myBoolean)

	const pi = 3.1415
	fmt.Println(pi)

	if num1 > 60 {
		fmt.Println("passed")
	} else if num1 > 20 && num2 > 20 {
		fmt.Println("passed 2")
	} else if num1 > 60 || num2 > 20 {
		fmt.Println("passed 3")
	} else {
		fmt.Println("Not passed")
	}

}

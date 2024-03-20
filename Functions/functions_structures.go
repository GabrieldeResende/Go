package main

import (
	"fmt" 
	"errors"
)

func main() {
	printMe("text from printMe function")
	var resultAdd = add(6,2)
	fmt.Println(resultAdd)
	var resultDivision, err = division(6,2)
	if err!=nil{
		fmt.Println(err.Error())
	}
	fmt.Println(resultDivision)

}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func add(num1 int, num2 int) int {
	var result = num1 + num2
	return result
}

func division(num1 int, num2 int) (int, error) {
	var err error
	if num2 == 0 {
		err = errors.New("Cant divide by 0")
		return 0, err
	}
	var result = num1 / num2
	return result, err
}

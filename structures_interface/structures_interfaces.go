package main

import "fmt"

type gasEngine struct {
	mpg uint8
	gallons uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("Vai da bom")
	} else {
		fmt.Println("Não vai da bom")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 15}
	fmt.Println(myEngine.milesLeft())
	canMakeIt(myEngine, 120)
}
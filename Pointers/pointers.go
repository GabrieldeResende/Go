package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	var i int32

	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(i)
}
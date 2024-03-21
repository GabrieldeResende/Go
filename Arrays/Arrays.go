package main

import "fmt"

func main() {
	var intArr [3]int32
	intArr[0] = 2
	fmt.Println(intArr[0])

	var intSlice []int32 = []int32{4,5,6}
	fmt.Println(intSlice)
	fmt.Printf("O tamanho é %v com sua capacidade %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7) //acrescentar elemento
	fmt.Println("\n", intSlice)
	fmt.Println("\n", intSlice, intArr)
	fmt.Printf("O tamanho é %v com sua capacidade %v", len(intSlice), cap(intSlice))
	
	var myMap map[string]uint8 = make(map[string]uint8)
	myMap = map[string]uint8{"Gabriel": 25}
	fmt.Println(myMap["Gabriel"])

	var i int = 0

	for i <= 10 {
		fmt.Println(i)
		i += 1
	}

	for j:= 0; j<=10; j++ {
		fmt.Println(j)
	}
}
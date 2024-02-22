package main

import "fmt"

func main() {
	array := []string{
		"zero",
		"one",
		"two",
	}

	fmt.Println("that is loop with index")
	for index := range array {
		fmt.Println(index)
		fmt.Println(array[index])
	}

	fmt.Println("that is loop without index")
	for _, noIndex := range array {
		fmt.Println(noIndex)
		// fmt.Println(array[noIndex]) // не будет работать тк коплилятор даже не пытается запоминать индекс
	}
}

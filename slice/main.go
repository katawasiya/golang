package main

import "fmt"

func main() {
	slice := make([]int, 0, 3)

	slice = append(slice, 1)
	slice = append(slice, 2)

	addElement(slice)

	fmt.Println("this is slice native slice: ", slice)
}

func addElement(slc []int) {
	slc = append(slc, 4)
	fmt.Println("this is addElement: ", slc)
	fmt.Println("len inside addElement: ", len(slc)) // Но изменение длины среза не отразится в вызывающей функции
}

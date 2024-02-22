package main

import "fmt"

func insertAds(freeItems []int64, paidItems []int64) ([]int64, error) {
	result := make([]int64, len(freeItems)+len(paidItems))

	j := 0
	for i := 0; i < len(result); i++ {
		if i%3 == 0 && j < len(paidItems) {
			result[i] = paidItems[j]
			j++
		} else {
			result[i] = freeItems[i-j]
		}
	}

	return result, nil
}

func main() {
	freeItems := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	paidItems := []int64{17, 18, 19, 20, 21, 22}

	result, err := insertAds(freeItems, paidItems)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(result)
}

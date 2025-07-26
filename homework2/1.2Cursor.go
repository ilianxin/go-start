package main

import "fmt"

func handleNums(nums *[]int) {
	for i := 0; i < len(*nums); i++ {
		(*nums)[i] = (*nums)[i] * 2
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	handleNums(&nums)
	fmt.Println(nums)

}

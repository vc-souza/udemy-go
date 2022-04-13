package main

import "fmt"

func main() {
	nums := []int{}

	for i := 0; i <= 10; i++ {
		nums = append(nums, i)
	}

	for _, num := range nums {
		msg := fmt.Sprintf("%d is ", num)

		if num%2 == 0 {
			msg += "even"
		} else {
			msg += "odd"
		}

		fmt.Println(msg)
	}
}

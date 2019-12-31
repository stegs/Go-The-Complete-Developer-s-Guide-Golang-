package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range nums {
		if math.Mod(float64(nums[i]), 2) == 0 {
			fmt.Println(i, "even")
		}

		if math.Mod(float64(nums[i]), 2) != 0 {
			fmt.Println(i, "odd")
		}
	}
}

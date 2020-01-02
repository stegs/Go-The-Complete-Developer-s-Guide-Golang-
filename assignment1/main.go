package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	/*for i := range nums {
		if math.Mod(float64(nums[i]), 2) == 0 {
			fmt.Println(i, "even")
		}

		if math.Mod(float64(nums[i]), 2) != 0 {
			fmt.Println(i, "odd")
		}
	}*/

	for _, num := range nums {
		if num%2 == 0 {
			fmt.Println(num, "even")
		} else if num%2 != 0 {
			fmt.Println(num, "odd")
		}
	}
}

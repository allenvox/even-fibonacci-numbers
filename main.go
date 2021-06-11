package main

import "fmt"

func main() {
	sum := 2             // because the first even Fibonacci number is 2
	nums := [2]int{1, 2} // buffer for ongoing neighbor-numbers
	i := 1
	for i < 4000000 {
		i = nums[0] + nums[1]
		if i%2 == 0 && i < 4000000 {
			fmt.Println(i)
			sum += i
		}
		nums[0] = nums[1]
		nums[1] = i
	}
	fmt.Println(sum)
}

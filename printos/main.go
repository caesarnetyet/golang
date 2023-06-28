package main

import "fmt"

func main() {
	slices := []int{3, 4, 5, 6}

	fmt.Println(slices[1+1:], slices[:1])
}

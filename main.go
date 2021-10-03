package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	i := 6
	copy(a[i:], a[i+1:]) // Shift a[i+1:] left one index.
	a[len(a)-1] = 0      // Erase last element (write zero value).
	a = a[:len(a)-1]
	fmt.Println(a)

}

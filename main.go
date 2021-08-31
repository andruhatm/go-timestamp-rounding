package main

// Importing fmt and time
import "fmt"

func main() {

	t := 1628729200000

	//timestamp const of needed round value (15 min)
	dev := 900000

	r := t % dev
	var fiftenStamp int
	if r != 0 {
		fiftenStamp = t + (dev - r)
	} else {
		fiftenStamp = t
	}

	// Prints output
	fmt.Printf("The result after rounding 't' is: %v\n", fiftenStamp)
}

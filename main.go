package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ranswer() int {
	rand.Seed(int64(time.Now().Day()))
	return rand.Intn(2)
}

func main() {
	var x string

	if ranswer() == 0 {
		x = "No"
	} else {
		x = "Yes"
	}
	fmt.Println(time.Now().Date())
	fmt.Println("Is Roman going to lunch today?")
	fmt.Println(x)
}
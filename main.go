package main

import (
	"fmt"
	"math/rand"
	"time"
	"flag"
)

func randomNumber(rmode bool) int {
	var seedVar int64

	if rmode == false {
		seedVar = int64(time.Now().Day())
	} else {
		seedVar = time.Now().Unix()
	}

	rand.Seed(seedVar)
	return rand.Intn(2)
}

func main() {
	var romanMode *bool
	var answer string

	romanMode = flag.Bool("absolutely-random", false, "Use absolute randomness mode (a bool)")
	flag.Parse()

	if randomNumber(*romanMode) == 0 {
		answer = "No"
	} else {
		answer = "Yes"
	}

	fmt.Println(time.Now().Date())
	fmt.Println("Is Roman going to lunch today?")
	fmt.Println(answer)
}
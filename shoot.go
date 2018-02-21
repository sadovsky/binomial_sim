package main

import (
	"fmt"
	"math/rand"
	"sort"
)

var trials [100000]int
var hit bool
var shots int

func main() {
	for i := 0; i < len(trials); i++ {
		hit = false
		shots = 0
		for hit == false {
			shots++
			if rand.Intn(10) == 0 {
				hit = true
				trials[i] = shots
			}
		}
	}

	//shit, go doesn't have a median function so lets hack one
	sort.Ints(trials[:])
	//even number of elements... but meh, this is close enough
	fmt.Println(trials[50000])
}

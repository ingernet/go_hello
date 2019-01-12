package main

import (
	"fmt"
	"math/rand"
	"time"
)

// populate range of length of English alphabet
var alphanums [26]int
var randalpha int

func toChar(i int) rune {
	return rune('A' - 1 + i)
}

func returnRandomAlpha() rune {
	for i := 1; i <= 26; i++ {
		alphanums[i-1] = i
	}
	randalpha := alphanums[rand.Intn(len(alphanums)-1)]
	fmt.Printf("%q\n", toChar(randalpha))
	return toChar(randalpha)
}

func main() {

	// print a random letter of the alphabet
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Printf("hello again, world\n")
	fmt.Printf("printing 10 letters, chosen at random:\n")

	for j := 1; j <= 10; j++ {
		returnRandomAlpha()
	}
}

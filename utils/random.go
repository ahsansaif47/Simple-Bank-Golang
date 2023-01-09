package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabets = "abcdefghijklmnopqrstuvwxys"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Create a random integer within a range..
func RandomInt(min, max int64) int64 {
	// rand.Int63n returns a number from 0 -> n
	// we add min to start it from minimum value provided by user..

	/*
		1. 	0 + min becomes min
		2. 	rand.Int63n(max - min + 1) + min returns max
			as it generates value 1 less than maximun limit.
	*/
	return min + rand.Int63n(max-min+1)
}

// Create a random string of length 'n'..
func RandomString(n int) string {
	var sb strings.Builder
	t_alphas := len(alphabets)

	// Generating n lengthed random string
	for i := 0; i < n; i++ {
		// picking a random characters from 26 alphas.
		c := alphabets[rand.Intn(t_alphas)]
		// appending the character to string builder buffer..
		sb.WriteByte(c)
	}
	// type-casting string builder format to strings..
	return sb.String()
}

// Creates random owner names using random strings..
func RandomOwner() string {
	return RandomString(6)
}

// Random money between 0 -> 1000
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// Pick a random currency from the mentioned currencies..
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	t_curr := len(currencies)
	return currencies[rand.Intn(t_curr)]
}

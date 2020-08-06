package random

import (
	"math/rand"
	"time"
)

//GenerateRandomString : Generate a random string
func GenerateRandomString() string {
	//global variable
	var (
		digits    string // sample of allowed number for random string
		specials  string // sample of allowed special character for random string
		alphabets string // sample of allowed alphabet for random string
		allowed   string // sample of all allowed character/digit for random string
		length    int    // the length of generated random string
	)

	//initialize value
	digits = "0123456789"
	specials = "~=+%^*/()[]{}/!@#$?|"
	alphabets = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyz" // all uppercase and lowercase alphabet
	allowed = alphabets + digits + specials
	length = 10 // will generate a random string with 10 characters long

	//prepare the random string behavior. this time, we use the uniqueness of unix timestamp
	//with timestamp behavior that change every second, our generated string will be constantly changing
	rand.Seed(time.Now().UnixNano())

	//prepare the buffer
	buf := make([]byte, length)

	//generate the random digits
	buf[0] = digits[rand.Intn(len(digits))]

	//generate the random special character
	buf[1] = specials[rand.Intn(len(specials))]

	//combined with all allowed character/symbol
	for i := 2; i < length; i++ {
		buf[i] = allowed[rand.Intn(len(allowed))]
	}

	//shuffle the combined above
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})

	return string(buf)

}

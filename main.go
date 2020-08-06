package main

import (
	"log"
	rnd "randomstring/random"
)

func main() {
	//call the method to generate a random string
	randomString := rnd.GenerateRandomString()
	log.Println(randomString)
}

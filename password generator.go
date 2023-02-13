package main

import (
	"fmt"
	"math/rand"
	"time"
)

const alphanum = "0123456789~!@#$%^&*()<>{}+-=?ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func main() {
	//the number of password digits
	length := 12

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		fmt.Printf("%c", alphanum[rand.Intn(len(alphanum))])
	}
}

package main

import (
	"fmt"
	"math/rand"
	"os"
)

func generatePassword() string {
	var chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()_+{}[]|\\\\;:'\\\",./<>?\"" //chars for our password
	var password string

	for i := 0; i <= 36; i++ {
		password += string(chars[rand.Intn(len(chars))]) // generate random string with loop for, that range is 36
	}
	return password
}
func main() {

	file, err := os.Create("passwords.txt") // creating txt file
	if err != nil {
		fmt.Printf("couldn't write: ", err)
	}
	for i := 0; i <= 6; i++ { // loop for to return our generated password and print it to the txt file
		password := generatePassword()
		fmt.Println(password)
		fmt.Fprintln(file, password)
	}
}

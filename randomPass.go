package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/fatih/color"
)

const (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits  = "0123456789"
	symbols = "!@#$%^&*_+=-|:;,./"
)

var passwordLength int
var useDigitsAnswer, specialSymbolsAnswer, forWhat string
var positiveResponse string = "yes"

func generatePassword(length int, useDigits, specialSymbols string) string {
	var runes []rune

	runes = append(runes, []rune(letters)...)
	if useDigitsAnswer == positiveResponse {
		runes = append(runes, []rune(digits)...)
	}
	if specialSymbolsAnswer == positiveResponse {
		runes = append(runes, []rune(symbols)...)
	}

	rand.NewSource(time.Now().UnixNano())

	password := make([]rune, length)
	for i := range password {
		password[i] = runes[rand.Intn(len(runes))]
	}

	return string(password)
}

func main() {

	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	passFile, createError := os.OpenFile("passwords.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if createError != nil {
		fmt.Println(red("Error - "))
		return
	}
	defer passFile.Close()

	fmt.Println(green("What do you want to create a password for?"))
	fmt.Scan(&forWhat)

	fmt.Println(green("Write password length"))
	fmt.Scan(&passwordLength)

	if passwordLength < 8 {
		fmt.Println(red("Password must contain at least 8 characters.Try Again"))
		os.Exit(0)
	}

	fmt.Println(green("If you want to use numbers in your password - write <<yes>> or <<no>>"))
	fmt.Scan(&useDigitsAnswer)

	fmt.Println(green("Can i use special symbols?  (yes or no)"))
	fmt.Scan(&specialSymbolsAnswer)

	if useDigitsAnswer != positiveResponse && specialSymbolsAnswer != positiveResponse {
		fmt.Println(red("I can't make the password for you with your answers("))
	}

	length := passwordLength
	useDigits := useDigitsAnswer
	specialSymbols := specialSymbolsAnswer

	password := generatePassword(length, useDigits, specialSymbols)

	passFile.WriteString(forWhat + " : " + password + "\n")

	fmt.Println(green("Your password is:"), password, green("And saved in passwords.txt file"))
}

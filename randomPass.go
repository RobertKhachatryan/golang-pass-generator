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
var useDigitsAnswer string
var useLettersAnswer string
var specialSymbolsAnswer string

func generatePassword(length int, useDigits, useLetters, specialSymbols string) string {
	var runes []rune

	positiveResponse := "yes"

	if useDigitsAnswer == positiveResponse {
		runes = append(runes, []rune(digits)...)
	}
	if useLettersAnswer == positiveResponse {
		runes = append(runes, []rune(letters)...)
	}
	if specialSymbolsAnswer == positiveResponse {
		runes = append(runes, []rune(symbols)...)
	}
	if useDigitsAnswer != positiveResponse && useLettersAnswer != positiveResponse && specialSymbolsAnswer != positiveResponse {
		fmt.Println("What the f*ck do you want !?")
	}

	rand.Seed(time.Now().UnixNano())

	password := make([]rune, length)
	for i := range password {
		password[i] = runes[rand.Intn(len(runes))]
	}

	return string(password)
}

func main() {

	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	fmt.Println(green("Write password length"))
	fmt.Scan(&passwordLength)

	if passwordLength < 8 {
		fmt.Println(red("Password must be contains 8 symbols.Try Again"))
		os.Exit(0)
	}

	fmt.Println(green("if you want to use digits in your password - write <<yes>> or <<no>>"))
	fmt.Scan(&useDigitsAnswer)

	fmt.Println(green("And if you want to use letters in your password - write <<yes>> or <<no>>"))
	fmt.Scan(&useLettersAnswer)

	fmt.Println(green("Can i use special symbols?  (yes or no)"))
	fmt.Scan(&specialSymbolsAnswer)

	length := passwordLength
	useDigits := useDigitsAnswer
	useLetters := useLettersAnswer
	specialSymbols := specialSymbolsAnswer

	password := generatePassword(length, useDigits, useLetters, specialSymbols)
	fmt.Println("Your password:", password)
}

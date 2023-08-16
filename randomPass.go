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
var useDigitsAnswer, useLettersAnswer, specialSymbolsAnswer, forWhat string
var positiveResponse string = "yes"

func generatePassword(length int, useDigits, useLetters, specialSymbols string) string {
	var runes []rune

	if useDigitsAnswer == positiveResponse {
		runes = append(runes, []rune(digits)...)
	}
	if useLettersAnswer == positiveResponse {
		runes = append(runes, []rune(letters)...)
	}
	if specialSymbolsAnswer == positiveResponse {
		runes = append(runes, []rune(symbols)...)
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

	passFile, createError := os.OpenFile("passwords.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if createError != nil {
		fmt.Println(red("Error - ", createError))
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

	fmt.Println(green("And if you want to use letters in your password - write <<yes>> or <<no>>"))
	fmt.Scan(&useLettersAnswer)

	fmt.Println(green("Can i use special symbols?  (yes or no)"))
	fmt.Scan(&specialSymbolsAnswer)

	if useDigitsAnswer != positiveResponse && useLettersAnswer != positiveResponse && specialSymbolsAnswer != positiveResponse {
		fmt.Println(red("I can't make the password for you with your answers("))
	}

	length := passwordLength
	useDigits := useDigitsAnswer
	useLetters := useLettersAnswer
	specialSymbols := specialSymbolsAnswer

	password := generatePassword(length, useDigits, useLetters, specialSymbols)

	passFile.WriteString(forWhat + " : " + password + "\n")

	fmt.Println(green("Your password is:"), password, green("And saved in passwords.txt file"))
}

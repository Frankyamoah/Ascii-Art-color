package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Color string

const (
	Black  = "\u001b[30m"
	Red    = "\u001b[31m"
	Green  = "\u001b[32m"
	Yellow = "\u001b[33m"
	Blue   = "\u001b[34m"
	Reset  = "\u001b[0m"
)

func main() {
	text := strings.Split(os.Args[1], "")

	one := Ascii(text, "standard")
	//fmt.Println(one)

	useColor := flag.String("color", Reset, "display colorized output")
	flag.Parse()

	if *useColor == Reset {
		colorize(Color(Green), "\r"+one)
		return
	}
	fmt.Println(one)
}

func colorize(color Color, message string) {
	fmt.Println(string(color), message, string(Reset))
}

func Ascii(text []string, font string) string {

	text2 := strings.Join(text, "")
	for _, word := range text2 {
		if word > 128 {
			return ""
		}
	}

	bytes, err := os.ReadFile(font + ".txt")
	if err != nil {
		fmt.Println(err)
		return ""
	}

	var newAscii []string
	if font == "thinkertoy" {
		newAscii = strings.Split(string(bytes), "\r\n")

	} else {
		newAscii = strings.Split(string(bytes), "\n")

	}

	var userInput []string
	var answer string
	newline := false

	for i := 0; i < len(text); i++ {

		if text[i] == "\\" && text[i+1] == "n" {
			newline = true
			continue
		}
		if newline {
			newline = false
			answer += printArtAscii(userInput, newAscii)
			userInput = []string{}

			continue

		}
		if text[i] == "\\" && text[len(text)-2] == "\\" {
			remove(text)
			continue

		}
		userInput = append(userInput, text[i])

	}

	answer += printArtAscii(userInput, newAscii)
	return answer
}

func printArtAscii(userInput []string, Ascii []string) string {
	empty := ""
	for line := 1; line <= 8; line++ {
		for _, word := range userInput {
			for _, character := range word {
				skip := (character - 32) * 9
				empty += Ascii[line+int(skip)]
			}
		}
		empty += "\n"
	}
	return empty
}

func remove(input []string) []string {
	var empty []string
	for _, character := range input {
		if character == "\\" {
			newstr := fmt.Sprint(input)
			New := strings.Replace(newstr, "\\", "", -1)
			empty = append(empty, New)
		}
	}
	return empty
}

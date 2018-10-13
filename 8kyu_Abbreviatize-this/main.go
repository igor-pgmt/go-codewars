// http://www.codewars.com/kata/abbreviatize-this/train/go
//
// We want to create a function Abbreviatize that takes a single parameter of type string, which returns the abbreviated form of the input string.

// Keep in mind that we want to make a shorter version of the word, so if the word is The, there's no point in abbreviating it to T1e. In that case, we just want to return the word that was passed in.

// Some examples:

// Abbreviatize("Globalization") => "G11n"
// Abbreviatize("Scherpenhuizen") => "S12n"
// Abbreviatize("Four") => "F2r"
// Abbreviatize("The") => "The"
// Abbreviatize("A") => "A"

package main

import (
	"fmt"
	"strconv"
)

func main() {

}

func Abbreviatize1(word string) string {
	if len(word) > 3 {
		return string(word[0]) + fmt.Sprint(len(word)-2) + string(word[len(word)-1])
	}
	return word
}

func Abbreviatize2(word string) string {
	x := len(word)
	if x > 3 {
		return fmt.Sprintf("%c%v%c", word[0], x-2, word[x-1])
	}
	return word
}

func Abbreviatize3(word string) string {
	if len(word) < 4 {
		return word
	}
	return fmt.Sprintf("%v%v%v", word[:1], len(word)-2, word[len(word)-1:])
}

func Abbreviatize4(word string) string {
	if len(word) <= 3 {
		return word
	}
	return string(word[0]) + strconv.Itoa(len(word[1:len(word)-1])) + string(word[len(word)-1])
}

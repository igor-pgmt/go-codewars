// https://www.codewars.com/kata/517abf86da9663f1d2000003
// Convert string to camel case
// Complete the method/function so that it converts dash/underscore delimited words into camel casing. The first word within the output should be capitalized only if the original word was capitalized.
// Examples
// ToCamelCase("the-stealth-warrior"); // returns "theStealthWarrior"
// ToCamelCase("The_Stealth_Warrior"); // returns "TheStealthWarrior"

package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(ToCamelCase("the-stealth-warrior")) // "theStealthWarrior"
	fmt.Println(ToCamelCase("The_Stealth_Warrior")) // "TheStealthWarrior"

}

func ToCamelCase(s string) (result string) {
	words := strings.FieldsFunc(s, func(c rune) bool {
		return c == '-' || c == '_'
	})
	if len(words) > 0 {
		result += words[0]
		for i := 1; i < len(words); i++ {
			result += strings.Title(words[i])
		}
	}
	return
}

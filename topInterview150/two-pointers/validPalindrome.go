package main

import "unicode"

func isPalindrome(s string) bool {
	var str string = convert(s)
	j := len(str) - 1
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[j] {
			return false
		}
		j--
	}
	return true
}

func convert(s string) string {
	var newString string
	for _, ch := range s {
		if unicode.IsUpper(ch) {
			ch = unicode.ToLower(ch)
		}
		if unicode.IsDigit(ch) {
			newString += string(ch)
			continue
		}
		if !unicode.IsLetter(ch) {
			continue
		}
		newString += string(ch)
	}
	return newString
}

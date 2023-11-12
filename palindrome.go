package main

import "strings"

func isPalindrome(s string) bool {
	// Remove spaces and convert to lowercase for a simple comparison
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)

	// Check if the string is the same when read backward
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

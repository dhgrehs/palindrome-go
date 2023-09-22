package main

import (
	"fmt"

	"palindrome-go/palindrome"
)

// A palindrome is a string of characters that read the same forwards as they do backwards.
// For example, "madam" is a palindrome.
// Also, we can check if a number is a palindrome or not.
// For example, 121 is a palindrome, but 123 is not.
func main() {
	fmt.Printf("Is madam a palindrome? %v\n", palindrome.IsStringPalindrome4("madam"))
	fmt.Printf("Is 121 a palindrome? %v\n", palindrome.IsIntegerPalindrome2(121))
	fmt.Printf("Is 123 a palindrome? %v\n", palindrome.IsIntegerPalindrome2(123))
}

package palindrome

import "strconv"

// Leverage Double Index via String  — O(?) Time, O(N) Space
func IsIntegerPalindrome1(number int) bool {
	str := strconv.Itoa(number)
	//We got a pretty efficient string palindrome algorithm,
	//so why not just convert the number to a string and use that?
	//The conversion is very expensive.
	return IsStringPalindrome4(str)
}

// Divide and Conquer  — O(N) Time, O(1) Space
func IsIntegerPalindrome2(number int) bool {
	originalNumber := number
	reverseNumber := 0

	//Processor are optimized to do math operations, so we can use that to our advantage.
	for number > 0 {
		// get unit's digit
		var remainder = number % 10
		// add unit's digit to bottom of reverseNumber
		reverseNumber = (reverseNumber * 10) + remainder
		// remove unit's digit from input number
		number = number / 10
	}

	return originalNumber == reverseNumber
}

package palindrome

// String Concatenation — O(N²) Time, O(N) Space
func IsStringPalindrome1(str string) bool {
	//First, we initialize an empty string we cal reversedStr
	reversedStr := ""

	//We'll iterate over the source string from the end to the beginning and append each character to reversedStr.
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}

	//Now that we have our reversedStr, we must compare char by char if they are equal or not.
	for i := range str {
		//Diff character means failure
		if str[i] != reversedStr[i] {
			return false
		}
	}
	//If we already checked all chars and did not find a mismatch, we can return true (is palindrome)
	return true
}

// Array Appending — O(N) Time, O(N) Space
func IsStringPalindrome2(str string) bool {
	//This approach will leverage the append usage, which is generally a constant time operation.
	res := []byte{}
	for i := len(str) - 1; i >= 0; i-- {
		res = append(res, str[i])
	}

	return str == string(res)
}

// Opposite Indices — O(N) Time, O(1) Space
func IsStringPalindrome3(str string) bool {
	//Instead of building the string in reverse order,
	//why not just check if the very pattern of the palindrome exists in the string?

	for i := 0; i < len(str); i++ {
		j := len(str) - 1 - i
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

// Opposite Indices w/ Half Check — O(N) Time, O(1) Space
func IsStringPalindrome4(str string) bool {
	//Instead of checking the entire string, since the first half should match the second,
	//we can just check the first half and compare it to the second half.
	if len(str) == 0 {
		return true
	}

	for i := 0; i < (len(str)/2)+1; i++ {
		j := len(str) - 1 - i
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

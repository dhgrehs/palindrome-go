package palindrome

import (
	"fmt"
	"strings"
	"testing"
)

// go test -run TestFunctionalInteger
// PASS
// ok  	palindrome-go/palindrome	0.218s
func TestFunctionalInteger(t *testing.T) {

	var palindromeFuncs = []func(int) bool{
		IsIntegerPalindrome1,
		IsIntegerPalindrome2}

	var tests = []struct {
		number int
		want   bool
	}{
		{999999999, true},
		{0, true},
		{123321, true},
		{1, true},
		{989999999, false},
	}
	for _, tt := range tests {

		for _, f := range palindromeFuncs {
			funcName := strings.Split(GetFunctionName(f), ".")
			testname := fmt.Sprintf("%v_%v", funcName[len(funcName)-1], tt.number)

			t.Run(testname, func(t *testing.T) {
				ans := f(tt.number)
				if ans != tt.want {
					t.Errorf("got %v, want %v", ans, tt.want)
				}
			})
		}

	}
}

// go test -bench=BenchmarkInteger -benchmem -benchtime=3s -count 3
// goos: darwin
// goarch: arm64
// pkg: palindrome-go/palindrome
// BenchmarkInteger/IsIntegerPalindrome1-8         	144016131	        24.90 ns/op	      16 B/op	       1 allocs/op
// BenchmarkInteger/IsIntegerPalindrome1-8         	144443737	        24.93 ns/op	      16 B/op	       1 allocs/op
// BenchmarkInteger/IsIntegerPalindrome1-8         	144351746	        25.02 ns/op	      16 B/op	       1 allocs/op
// BenchmarkInteger/IsIntegerPalindrome2-8         	298039315	        12.11 ns/op	       0 B/op	       0 allocs/op
// BenchmarkInteger/IsIntegerPalindrome2-8         	296581567	        12.04 ns/op	       0 B/op	       0 allocs/op
// BenchmarkInteger/IsIntegerPalindrome2-8         	301276921	        12.14 ns/op	       0 B/op	       0 allocs/op
// PASS
// ok  	palindrome-go/palindrome	32.852s
func BenchmarkInteger(b *testing.B) {
	var palindromeFuncs = []func(int) bool{
		IsIntegerPalindrome1,
		IsIntegerPalindrome2}

	number := 99999999999
	for _, f := range palindromeFuncs {

		funcName := strings.Split(GetFunctionName(f), ".")
		testname := fmt.Sprintf("%v", funcName[len(funcName)-1])

		b.Run(testname, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				f(number)
			}
		})
	}

}

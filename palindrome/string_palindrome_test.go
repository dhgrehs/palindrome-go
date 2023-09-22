package palindrome

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

// go test -run TestFunctionalString
// PASS
// ok  	palindrome-go/palindrome	0.483s
func TestFunctionalString(t *testing.T) {

	var palindromeFuncs = []func(string) bool{
		IsStringPalindrome1,
		IsStringPalindrome2,
		IsStringPalindrome3,
		IsStringPalindrome4}

	var tests = []struct {
		str  string
		want bool
	}{
		{"saippuakivikauppias", true},
		{"", true},
		{"madam", true},
		{"not", false},
		{"asbtyhjk", false},
	}
	for _, tt := range tests {

		for _, f := range palindromeFuncs {
			funcName := strings.Split(GetFunctionName(f), ".")
			testname := fmt.Sprintf("%v_%v", funcName[len(funcName)-1], tt.str)

			t.Run(testname, func(t *testing.T) {
				ans := f(tt.str)
				if ans != tt.want {
					t.Errorf("got %v, want %v", ans, tt.want)
				}
			})
		}

	}
}

// go test -bench=BenchmarkString -benchmem -benchtime=3s -count 3
//goos: darwin
//goarch: arm64
//pkg: palindrome-go/palindrome
//BenchmarkString/IsStringPalindrome1-8         	 6556833	       530.2 ns/op	     312 B/op	      37 allocs/op
//BenchmarkString/IsStringPalindrome1-8         	 6751858	       536.2 ns/op	     312 B/op	      37 allocs/op
//BenchmarkString/IsStringPalindrome1-8         	 6750886	       536.1 ns/op	     312 B/op	      37 allocs/op
//BenchmarkString/IsStringPalindrome2-8         	65722164	        54.87 ns/op	      56 B/op	       3 allocs/op
//BenchmarkString/IsStringPalindrome2-8         	65713665	        55.03 ns/op	      56 B/op	       3 allocs/op
//BenchmarkString/IsStringPalindrome2-8         	65679897	        56.18 ns/op	      56 B/op	       3 allocs/op
//BenchmarkString/IsStringPalindrome3-8         	392423090	         9.163 ns/op	       0 B/op	       0 allocs/op
//BenchmarkString/IsStringPalindrome3-8         	392615858	         9.168 ns/op	       0 B/op	       0 allocs/op
//BenchmarkString/IsStringPalindrome3-8         	391612884	         9.170 ns/op	       0 B/op	       0 allocs/op
//BenchmarkString/IsStringPalindrome4-8         	486578952	         7.411 ns/op	       0 B/op	       0 allocs/op
//BenchmarkString/IsStringPalindrome4-8         	487987624	         7.472 ns/op	       0 B/op	       0 allocs/op
//BenchmarkString/IsStringPalindrome4-8         	486624637	         7.441 ns/op	       0 B/op	       0 allocs/op
//PASS
//ok  	palindrome-go/palindrome	50.732s

func BenchmarkString(b *testing.B) {
	var palindromeFuncs = []func(string) bool{
		IsStringPalindrome1,
		IsStringPalindrome2,
		IsStringPalindrome3,
		IsStringPalindrome4}

	word := "saippuakivikauppias"
	for _, f := range palindromeFuncs {

		funcName := strings.Split(GetFunctionName(f), ".")
		testname := fmt.Sprintf("%v", funcName[len(funcName)-1])

		b.Run(testname, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				f(word)
			}
		})
	}

}

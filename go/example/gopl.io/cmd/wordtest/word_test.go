package word

import (
	"fmt"
	"testing"
)

func TestPalindrome(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Error(`IsPalindrome("detartrated") = false`)
	}
	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") = false`)
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Error(`IsPalindrome("palindrome") = true`)
	}
}

func BenchmarkPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("detartrated")
	}
}

func ExampleIsPalindrome() {
	fmt.Println(IsPalindrome("detartrated"))
	fmt.Println(IsPalindrome("palindrome"))
	// output:
	// true
	// false
}

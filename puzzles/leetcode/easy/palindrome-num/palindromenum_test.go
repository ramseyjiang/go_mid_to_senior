package palindrome_num

import (
	"log"
	"testing"
)

func TestIsPalindrome1True(t *testing.T) {
	expected := isPalindrome1(123321)
	if expected == true {
		log.Println("int num is palindrome")
	} else {
		t.Error("int num is not palindrome")
	}
}

func TestIsPalindrome1False(t *testing.T) {
	expected := isPalindrome1(-121) // 10
	if expected == false {
		log.Println("int num is not palindrome")
	} else {
		t.Error("int num is palindrome")
	}
}

func TestIsPalindrome2True(t *testing.T) {
	expected := isPalindrome2(123321)
	if expected == true {
		log.Println("int num is palindrome")
	} else {
		t.Error("int num is not palindrome")
	}
}

func TestIsPalindrome2False(t *testing.T) {
	expected := isPalindrome2(-121) // 10
	if expected == false {
		log.Println("int num is not palindrome")
	} else {
		t.Error("int num is palindrome")
	}
}

package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	s := "12345678987654321"
	if !IsPalindrome(Palindrome(s)) {
		t.Fatalf("fail: %s\n", s)
	}

	s = "asdfgfdsa"
	if !IsPalindrome(Palindrome(s)) {
		t.Fatalf("fail: %s\n", s)
	}

	s = "あいうえおえういあ"
	if !IsPalindrome(Palindrome(s)) {
		t.Fatalf("fail: %s\n", s)
	}

	s = "abcあいう123"
	if IsPalindrome(Palindrome(s)) {
		t.Fatalf("fail: %s\n", s)
	}
}

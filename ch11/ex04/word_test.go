package word

import (
	"math/rand"
	"testing"
	"time"
	"unicode"
)

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i, j := 0, n-1; i < j; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		//
		if !unicode.IsSpace(r) && !unicode.IsPunct(r) {
			runes[j] = r
			j--
		}
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

func randomNonPalindrome(rng *rand.Rand) string {
	n := rng.Intn(23) + 2
	p := make([]rune, n)
	for i := 0; i < n; i++ {
		for {
			r1 := rune(rng.Intn(0x1000))
			r2 := rune(rng.Intn(0x1000))
			if unicode.IsLetter(r1) && unicode.IsLetter(r2) && unicode.ToLower(r1) != unicode.ToLower(r2) {
				p[i] = r1
				p[n-1-i] = r2
				break
			}
		}
	}
	return string(p)
}

func TestRandomNonPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomNonPalindrome(rng)
		if IsPalindrome(p) {
			t.Errorf("IsPalindrome(%v) = true", p)
		}
	}
}
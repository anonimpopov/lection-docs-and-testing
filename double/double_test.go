package double

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unicode/utf8"
)

func TestDouble(t *testing.T) {
	// Arrange
	cases := []struct {
		name   string
		input  int
		expect int
	}{
		{
			"first 1",
			1,
			2,
		},
		{
			"try",
			0,
			0,
		},
	}

	for _, cs := range cases {
		t.Run(cs.name, func(t *testing.T) {
			// Act
			res := Double(cs.input)

			// Assert
			assert.Equal(t, cs.expect, res)
		})
	}

}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}

func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}
	for _, tc := range testcases {
		rev := Reverse(tc.in)
		if rev != tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}
}

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev := Reverse(orig)
		doubleRev := Reverse(rev)
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}

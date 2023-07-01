package gohello

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestCrrrName(t *testing.T) {
	name := "Anton"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Crrr(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Crrr("Anton") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestCrrrEmpty(t *testing.T) {
	msg, err := Crrr("")
	if msg != "" || err == nil {
		t.Fatalf(`Crrr("") = %q, %v, want "", error`, msg, err)
	}
}

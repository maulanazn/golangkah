package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "maulana"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("maulana")
	if !want.MatchString(name) || err != nil {
		t.Fatalf(`Hello("maulana") = %q, %v, want match %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

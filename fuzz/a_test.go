package fuzz

import "testing"

func TestReverseString(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Muinde", "edniuM"},
		{"omo", "omo"},
	}

	for _, c := range cases {
		result := ReverseString(c.in)
		if result != c.want {
			t.Errorf("Passed (%q), Expected (%q), Got (%q)", c.in, result, c.want)
		}
	}
}

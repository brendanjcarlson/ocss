package types

import "testing"

func Test_IsNumber(t *testing.T) {
	testcases := []struct {
		input Literal
		want  bool
	}{
		{"", false},
		{"a", false},

		{"+", false},
		{"-", false},
		{"..", false},
		{"+.", false},
		{"-.", false},

		{"1.1.", false},
		{"+1.1.", false},
		{"-1.1.", false},

		{"1", true},
		{"+1", true},
		{"-1", true},

		{".1", true},
		{"+.1", true},
		{"-.1", true},

		{"1.0", true},
		{"+1.0", true},
		{"-1.0", true},
	}

	for _, testcase := range testcases {
		t.Run("is_number", func(t *testing.T) {
			got := IsNumber(testcase.input)
			if testcase.want != got {
				t.Fatalf("expected %t for literal %q, got %t", testcase.want, testcase.input, got)
			}
		})
	}
}

package types

import "testing"

func Test_IsPercentage(t *testing.T) {
	testcases := []struct {
		input Literal
		want  bool
	}{
		{"", false},
		{"a", false},

		{"1", false},
		{"1.0", false},

		{"+1", false},
		{"+1.0", false},
		{"-1", false},
		{"-1.0", false},

		{"+.1", false},
		{"-.1", false},

		{"1%", true},
		{"1.0%", true},

		{"+1%", true},
		{"-1%", true},

		{"+.1%", true},
		{"-.1%", true},
	}

	for _, testcase := range testcases {
		t.Run("is_percentage", func(t *testing.T) {
			got := IsPercentage(testcase.input)
			if testcase.want != got {
				t.Fatalf("expected %t for literal %q, got %t", testcase.want, testcase.input, got)
			}
		})
	}
}

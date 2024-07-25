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

		{"+1", false},
		{"-1", false},

		{"+.1", false},
		{"-.1", false},

		{"1%", true},

		{"+1%", true},
		{"-1%", true},

		{"+.1%", true},
		{"-.1%", true},
	}

	for _, testcase := range testcases {
		t.Run("is_percentage", func(t *testing.T) {
			got := testcase.input.IsPercentage()
			if testcase.want != got {
				t.Fatalf("expected %t for literal %q, got %t", testcase.want, testcase.input, got)
			}
		})
	}
}

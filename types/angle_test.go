package types

import "testing"

func Test_IsAngle(t *testing.T) {
	testcases := []struct {
		input Literal
		want  bool
	}{
		{"", false},
		{"1rem", false},

		{"1deg", true},
		{".1deg", true},
		{"1.1deg", true},
		{"-1.1deg", true},
		{"+1.1deg", true},

		{"1rad", true},
		{".1rad", true},
		{"1.1rad", true},
		{"-1.1rad", true},
		{"+1.1rad", true},

		{"1grad", true},
		{".1grad", true},
		{"1.1grad", true},
		{"-1.1turn", true},
		{"+1.1turn", true},

		{"1turn", true},
		{".1turn", true},
		{"1.1turn", true},
		{"-1.1turn", true},
		{"+1.1turn", true},
	}

	for _, testcase := range testcases {
		t.Run("is_angle", func(t *testing.T) {
			got := IsAngle(testcase.input)
			if testcase.want != got {
				t.Fatalf("expected %t for literal %q, got %t", testcase.want, testcase.input, got)
			}
		})
	}
}

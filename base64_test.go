package base64

import "testing"

type testCase struct {
	input  string
	output string
}

func (tc *testCase) run(t *testing.T) {
	d, err := encode(tc.input)
	if err != nil {
		t.Errorf("Unable to encode '%s'", err.Error())
	}
	if string(d) != tc.output {
		t.Errorf("Wrong encoding, expected '%s', actual '%s'", tc.output, string(d))
	}
}

func TestEncode(t *testing.T) {
	tests := []testCase{
		{
			input:  "Many hands make light work.",
			output: "TWFueSBoYW5kcyBtYWtlIGxpZ2h0IHdvcmsu",
		},
		{
			input:  "light wor",
			output: "bGlnaHQgd29y",
		},
		{
			input:  "light work.",
			output: "bGlnaHQgd29yay4=",
		},
		{
			input:  "light work",
			output: "bGlnaHQgd29yaw==",
		},
	}

	for _, test := range tests {
		test.run(t)
	}
}

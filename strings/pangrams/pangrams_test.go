package pangrams

import "testing"

//import "strings"

func TestPangram(t *testing.T) {
	cases := map[string]bool{
		"We promptly judged antique ivory buckles for the next prize": true,
		"We promptly judged antique ivory buckles for the prize":      false}
	for text, exp_res := range cases {
		res := Pangram(text)
		if res != exp_res {
			t.Errorf("%s got: %t, want: %t.", text, res, exp_res)
		}
	}
}

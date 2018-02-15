package pangrams

import (
	"fmt"
	"strings"
)

func Pangram(str string) bool {
	r := true
	s := ""
	fmt.Println(str)
	for i := 'a'; r && i <= 'z'; i++ {
		s = fmt.Sprintf("%q", i)
		r = strings.ContainsAny(str, s) || strings.ContainsAny(str, strings.ToUpper(s))
		//fmt.Printf("%q => %t %s\n", i, r, s)
	}
	return r
}

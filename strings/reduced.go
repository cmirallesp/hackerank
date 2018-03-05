package strings

//import (
//"fmt"
//)

func Reduced(str string) string {
	result := ""
	j := 0
	cur := ""
	r := ""
	for i := 0; i < len(str); i++ {
		cur = str[i : i+1]
		//fmt.Printf("1) %d => %s, %s %d\n", i, cur, r, j)
		if cur == r {
			result = str[0 : j-1]
			j -= 1
		} else {
			result += str[i : i+1]
			j += 1
		}
		//fmt.Printf("2) %d => %s\n", i, result)
		if j > 0 {
			r = result[j-1 : j]
		} else {
			r = ""
		}
	}
	//fmt.Printf("%s\n", result)
	if result == "" {
		return "Empty String"
	} else {
		return result
	}
}

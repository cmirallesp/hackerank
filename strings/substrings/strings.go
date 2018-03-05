package strings

import (
	"fmt"
	s "github.com/deckarep/golang-set"
)

func ForCountUniqueSubstrings(str string, begin, end int) int {
	total := 0
	all_substrings := make(map[string]bool)
	var ss string
	for j := end; j > begin-1; j-- {
		for i := j; i < end+1; i++ {
			ss = fmt.Sprintf("%s", str[j:i+1])
			//fmt.Printf("%s\n", ss)
			if _, ok := all_substrings[ss]; !ok {
				all_substrings[ss] = true
				total += 1
			}
		}
	}
	//fmt.Printf("%v\n", all_substrings)
	return total
}

func arrCountUniqueSubstrings(str string, i, j int, prev_substrings []string, all_substrings map[string]bool) int {
	if j < i {
		return len(all_substrings)
	} else {
		c := fmt.Sprintf("%s", str[j:j+1])
		my_substrings := make([]string, len(prev_substrings)+1)
		my_substrings[0] = c
		all_substrings[c] = true
		ss := ""
		for ii, v := range prev_substrings {
			ss = fmt.Sprintf("%s%s", c, v)
			my_substrings[ii+1] = ss
			all_substrings[ss] = true
		}
		return arrCountUniqueSubstrings(str, i, j-1, my_substrings, all_substrings)
	}
}

func ArrCountUniqueSubstrings(str string, i, j int) int {
	//a := []string{str[j : j+1]}
	return arrCountUniqueSubstrings(str, i, j, make([]string, 0), make(map[string]bool))
	//fmt.Printf("1) %s\n", result)
	//fmt.Printf("2) %s\n", r)
}

func countUniqueSubstrings(str string, i, j int) (s.Set, s.Set) {
	c := fmt.Sprintf("%s", str[i:i+1])
	if i == j {
		result := s.NewSet()
		result.Add(c)
		return result, result
	} else {
		next_prefixes, all_prefixes := countUniqueSubstrings(str, i+1, j)
		my_prefixes := s.NewSet()
		my_prefixes.Add(c)
		selem := ""
		it := next_prefixes.Iterator()
		for elem := range it.C {
			selem = fmt.Sprintf("%s", elem)
			s := fmt.Sprintf("%s%s", c, selem)
			my_prefixes.Add(s)
		}
		all_prefixes = all_prefixes.Union(my_prefixes)
		//fmt.Printf("my_prefixes=> %v all_prefixes => %v \n", my_prefixes, all_prefixes)
		return my_prefixes, all_prefixes
	}
}

func CountUniqueSubstrings(str string, i, j int) int {
	_, r := countUniqueSubstrings(str, i, j)
	return r.Cardinality()
	/*sl := r.ToSlice()*/
	//result := make([]string, len(sl))
	//for i, v := range sl {
	//result[i] = v.(string)
	//}
	//sort.Strings(result)
	/*return result*/
}

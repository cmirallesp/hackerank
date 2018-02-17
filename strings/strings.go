package strings

import (
	"fmt"
	s "github.com/deckarep/golang-set"
)

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

package utilstrings

import "strings"

//WordCount return the number of the words that contains in string
func WordCount(v string) int {
	words := strings.Fields(v)
	return len(words)
}

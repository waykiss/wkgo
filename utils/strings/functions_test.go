package utilstrings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWordCount(t *testing.T) {
	testCases := []struct {
		value    string
		expected int
	}{
		{"Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard", 19},
		{"", 0},
		{"rodrigo rodrigues", 2},
		{"rodrigo", 1},
		{" ", 0},
	}
	for _, v := range testCases {
		got := WordCount(v.value)
		assert.Equal(t, v.expected, got, "", v.value, v.expected, got)
	}
}

package arraysandhashmaps

import (
	"strings"
	"unicode"
)


func IsPalindromePermutation(str string) bool  {
	strLen := len(str)
	index := 0
	mirrorIndex := strLen -1
	for ;index < mirrorIndex;{
		char := str[index]
		mirrorChar := str[mirrorIndex]
		if !unicode.IsLetter(rune(char)){
			index++
			continue
		}

		if !unicode.IsLetter(rune(mirrorChar)){
			mirrorIndex--
			continue
		}

		if !strings.EqualFold(string(char), string(mirrorChar)){
			return false
		}
		index++
		mirrorIndex--
	}

	return true


}

package arraysandhashmaps

import (
	"strconv"
)

func CompressString(str string) (result string) {
	strLen := len(str)
	compressedString := ""
	result = str
	charRepeats := 0

	for i := 0; i < strLen; i++ {
		char := string(str[i])
		charRepeats++
		if i + 1 >= strLen || char != string(str[i+1]) {
			compressedString += char + strconv.Itoa(charRepeats)
			charRepeats = 0
		}
	}

	if len(compressedString) < strLen {
		result = compressedString
	}
	return result
}

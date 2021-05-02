package arraysandhashmaps

import (
	"strconv"
)

func CompressString(str string) string {
	strLen := len(str)
	compressedString := ""
	numberOfReplacements := 0

	var previousCharacter string
	repeatCount := 0
	for i := 0 ; i <= strLen ; i++{
		var char string
		if i< strLen{
			char = string(str[i])
		}

		repeatCount++
		if char == previousCharacter{
			continue
		}
		if i > 0{
			numberOfReplacements++
			compressedString += previousCharacter + strconv.Itoa(repeatCount)
		}
		repeatCount = 0
		previousCharacter = char
	}
	if (numberOfReplacements * 2) < strLen{
		return compressedString
	}
	return str

}
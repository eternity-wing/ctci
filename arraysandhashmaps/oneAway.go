package arraysandhashmaps


func IsOneAway(str1 string, str2 string) bool  {

	str1Index := 0
	str2Index := 0

	numberOfDiffs := 0

	str1Len := len(str1)
	str2Len := len(str2)


	if str2Len - str1Len > 1{
		return false
	}

	for ;str2Index < str2Len && numberOfDiffs <= 1 && str1Index < str1Len;{

		haveSameCharacterAtIndex := str1[str1Index] == str2[str2Index]
		if haveSameCharacterAtIndex{
			str1Index++
			str2Index++
			continue
		}
		if str1Len == str2Len{
			str2Index++
		}
		numberOfDiffs++
		str1Index++
	}

	return numberOfDiffs <= 1
}
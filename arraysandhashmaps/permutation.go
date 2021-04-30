package arraysandhashmaps

func Permutation(str1 string, str2 string) (result bool)  {
	str1Len := len(str1)
	str2Len := len(str2)

	var occurredCharacters = make(map[string]int)

	if str1Len != str2Len{
		return false
	}

	for i := 0 ; i < str1Len ; i++{
		chr1 := string(str1[i])
		chr2 := string(str2[i])
		if _, ok := occurredCharacters[chr1]; !ok {
			occurredCharacters[chr1] = 0
		}

		if _, ok := occurredCharacters[chr2]; !ok {
			occurredCharacters[chr2] = 0
		}

		occurredCharacters[chr1] += 1
		occurredCharacters[chr2] -= 1
	}

	for _, contradictionsCount := range occurredCharacters{
		if contradictionsCount != 0 {
			return false
		}
	}
	return true
}
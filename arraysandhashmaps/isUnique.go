package arraysandhashmaps


func IsUnique(str string) bool  {
	var occurredCharacters = make(map[string]bool)

	for _, asciCode := range str{
		character := string(asciCode)
		if _, ok := occurredCharacters[character]; ok {
			return false
		}
		occurredCharacters[character] = true
	}
	return true
}
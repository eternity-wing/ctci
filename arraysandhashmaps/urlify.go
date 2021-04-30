package arraysandhashmaps


const ENCODED_SPACE = "%20"
const RAW_SPACE = " "

func Urlify(rawString string, length int) string  {
	var urlifiedString = ""
	for i:=0 ; i<length ; i++{
		chr := string(rawString[i])
		if chr == RAW_SPACE {
			urlifiedString += ENCODED_SPACE
		}else{
			urlifiedString += string(rawString[i])
		}
	}
	return urlifiedString
}
package reverse

func Reverse(str string) string {
	runes := []rune(str)
	result := make([]rune, len(runes))
	for i := 0; i < len(runes); i++ {
		result[i] = runes[len(runes)-1-i]
	}
	return string(result)
}

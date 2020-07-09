package raindrops

import "strconv"

// Convert factor into raindrop sounds
func Convert(rain int) string {

	result := ""

	if rain%3 == 0 {
		result += "Pling"
	}

	if rain%5 == 0 {
		result += "Plang"
	}

	if rain%7 == 0 {
		result += "Plong"
	}

	if result == "" {
		result = strconv.Itoa(rain)
	}

	return result

}

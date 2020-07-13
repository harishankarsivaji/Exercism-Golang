package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(array []string) FreqMap {
	var results = make(chan FreqMap, len(array))
	for _, str := range array {
		go func(str string) {
			results <- Frequency(str)
		}(str)
	}

	var ret = FreqMap{}
	for range array {
		m := <-results
		for k, v := range m {
			ret[k] += v
		}
	}
	return ret
}

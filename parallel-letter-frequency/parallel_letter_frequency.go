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

func freqWrapper(s string, ch chan FreqMap) {
	ch <- Frequency(s)
}

// ConcurrentFrequency uses concurrency to count rune frequencies
func ConcurrentFrequency(stringList []string) FreqMap {
	ch := make(chan FreqMap)
	for _, s := range stringList {
		go freqWrapper(s, ch)
	}

	returnMap := FreqMap{}
	for i := 0; i < len(stringList); i++ {
		newMap := <-ch
		for k, v := range newMap {
			returnMap[k] += v
		}
	}
	return returnMap
}

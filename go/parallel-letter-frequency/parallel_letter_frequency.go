package letter

import (
	"sync"
)

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

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	var wg sync.WaitGroup
	wg.Add(len(l))
	c := make(chan FreqMap)

	go func() {
		wg.Wait()
		close(c)
	}()

	for _, s := range l {
		go func(s string, c chan FreqMap) {
			defer wg.Done()
			c <- Frequency(s)
		}(s, c)
	}

	result := FreqMap{}

	for f := range c {
		for k, v := range f {
			result[k] += v
		}
	}

	return result
}

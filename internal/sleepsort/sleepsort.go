package sleepsort

import (
	"sync"
	"time"
)

// canonicalDuration returns a duration that represents its position when sorting.
func canonicalDuration(num uint8) time.Duration {
	return time.Duration(num) * time.Millisecond
}

// sleepThenBuffer schedules and adds a number to a buffer.
func sleepThenBuffer(buffer chan uint8, buffering *sync.WaitGroup, num uint8) {
	defer buffering.Done()

	time.Sleep(canonicalDuration(num))

	buffer <- num
}

// formatBuffer converts a buffer into a slice.
func formatBuffer(buffer chan uint8) ([]uint8, bool) {
	formatted := make([]uint8, len(buffer))

	for i := range formatted {
		s, ok := <-buffer
		if !ok {
			return formatted, false
		}

		formatted[i] = s
	}

	return formatted, true
}

// Sort performs the Sleep Sort and returns a sorted copy of the array.
func Sort(nums []uint8) ([]uint8, bool) {
	sortBuffer := make(chan uint8, len(nums))
	defer close(sortBuffer)

	var sorting sync.WaitGroup
	sorting.Add(len(nums))

	for _, n := range nums {
		go sleepThenBuffer(sortBuffer, &sorting, n)
	}

	sorting.Wait()

	return formatBuffer(sortBuffer)
}

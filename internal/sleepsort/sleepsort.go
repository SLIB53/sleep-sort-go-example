package sleepsort

import (
	"sync"
	"time"
)

func assignDuration(num uint8) time.Duration {
	return time.Duration(num) * time.Millisecond
}

func sleepThenBuffer(buffer chan uint8, buffering *sync.WaitGroup, num uint8) {
	defer buffering.Done()

	time.Sleep(assignDuration(num))

	buffer <- num
}

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

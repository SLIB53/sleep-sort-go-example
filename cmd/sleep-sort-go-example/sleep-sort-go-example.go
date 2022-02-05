package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func sleepThenRecieve(buffer chan uint8, buffering *sync.WaitGroup, num uint8) {
	time.Sleep(time.Duration(num) * time.Millisecond)

	buffer <- num

	buffering.Done()
}

func formatBuffer(buffer chan uint8) ([]uint8, bool) {
	formatted := make([]uint8, len(buffer))

	for i := 0; i < len(formatted); i++ {
		s, ok := <-buffer
		if !ok {
			return formatted, false
		}

		formatted[i] = s
	}

	return formatted, true
}

func sleepSort(nums []uint8) ([]uint8, bool) {
	sortBuffer := make(chan uint8, len(nums))
	defer close(sortBuffer)

	var sorting sync.WaitGroup
	sorting.Add(len(nums))

	for _, n := range nums {
		go sleepThenRecieve(sortBuffer, &sorting, n)
	}

	sorting.Wait()

	return formatBuffer(sortBuffer)
}

func main() {
	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))

	nums := make([]uint8, 8)
	for i := range nums {
		nums[i] = uint8(randomGenerator.Intn(256))
	}

	sortedNums, ok := sleepSort(nums)

	if !ok {
		fmt.Println("uh-oh... sleep sort failed for the following input: ", nums)
		return
	}

	fmt.Println("original:", nums)
	fmt.Println("  sorted:", sortedNums)
}

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/slib53/sleep-sort-go-example/internal/sleepsort"
)

func main() {
	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))

	nums := make([]uint8, 8)
	for i := range nums {
		nums[i] = uint8(randomGenerator.Intn(256))
	}

	sortedNums, ok := sleepsort.Sort(nums)

	if !ok {
		fmt.Println("uh-oh... sleep sort failed for the following input: ", nums)
		return
	}

	fmt.Println("original:", nums)
	fmt.Println("  sorted:", sortedNums)
}

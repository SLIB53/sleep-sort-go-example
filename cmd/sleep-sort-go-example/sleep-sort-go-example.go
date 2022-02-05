package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/slib53/sleep-sort-go-example/internal/sleepsort"
)

var randomGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))

// makeOriginal generates the original array to be sorted.
func makeOriginal() []uint8 {
	nums := make([]uint8, 5)

	for i := range nums {
		nums[i] = uint8(randomGenerator.Intn(math.MaxUint8 + 1))
	}

	return nums
}

func main() {
	original := makeOriginal()

	sorted, ok := sleepsort.Sort(original)

	if !ok {
		fmt.Println("uh-oh... sleep sort failed for the following input: ", original)
		return
	}

	fmt.Println("original:", original)
	fmt.Println("  sorted:", sorted)
}

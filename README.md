# Sleep Sort Go Example

Example implementation of the [sleep sort][sleep-sort] using Go concurrency.

The program will randomly generate and print an array, and then print the sorted array:

```
original: [3 147 23 242 44]
  sorted: [3 23 44 147 242]
```

## Usage

To run the example, use `go run`:

```sh
go run cmd/sleep-sort-go-example/sleep-sort-go-example.go
```

[sleep-sort]: https://rosettacode.org/wiki/Sorting_algorithms/Sleep_sort

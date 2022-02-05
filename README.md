# Sleep Sort Go Example

Example implementation of the [sleep sort][sleep-sort] using Go concurrency.

The program will randomly generate and print an array, and then print the sorted array:

```
original: [219 226 225 102 182 255 245 36]
  sorted: [36 102 182 219 225 226 245 255]
```

## Usage

To run the example, use `go run`:

```sh
go run cmd/sleep-sort-go-example/sleep-sort-go-example.go
```

[sleep-sort]: https://rosettacode.org/wiki/Sorting_algorithms/Sleep_sort

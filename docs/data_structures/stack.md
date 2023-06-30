# Stack

Last in → First out. Push, peek, pop works very fast in `O(1)` time, which is good.
It's mainly when you need to work with the latest inserted value.

You can implement it simply like an array:
```go
s := make([]int, 0)        // create
s = append(s, 1)           // push
poppedValue := s[len(s)-1] // peek
s = s[:len(s)-1]           // pop
```

Or you can use my generic stack data structure. Source code: [src/data_structures/stack.go](../../src/data_structures/stack.go)

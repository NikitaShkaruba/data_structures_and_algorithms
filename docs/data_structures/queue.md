# Queue

Last in → Last out. Enqueue, peek, dequeue work very fast in `O(1)` time, which is good.
It can be used when you need to process data in the order it was given to you.

You can implement it simply with an array. Simple implementation works, but has a memory leak in dequeue.
```go
q := make([]int, 0)        // create
q = append(q, 100)         // enqueue
queuedValue := q[len(q)-1] // peek
q = q[1:]                  // dequeue
```

Or you can use my generic queue data structure that uses linked list underneath. Source code: [src/data_structures/queue.go](../../src/data_structures/queue.go)

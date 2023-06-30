# Hashing

### Map

Map is the best data-structure ever. It has `O(1)` time complexity for reading, writing and deleting an element.
It's being used in `> 70%` of leetcode solutions. Try it everywhere :)

```go
m := make(map[string]int) // create

m["test"] = 2 // write

readValue := m["test"] // read

if readValue2, contains := m["test"]; contains {
	// logic
} // read + contains 

delete(m, "test") // delete

///// More create /////

m2 := map[string]int{"one": 1, "two": 2} // create 2

var m3 map[int]int = make(map[int]int) // create 3

s := "test"
m4 := map[byte]bool{
    s[0]: true, // dynamic initialization is fine
} // create 4
```

### Set

Set is just a map with no values, only keys. It's time and space complexities are the same.
We use empty struct as a value, because go shares values between each key.

```go
s := make(map[string]struct{}) // create

s["test"] = struct{}{} // write

if _, contains := s["test"]; contains {
	// logic
} // contains

delete(s, "test") // delete
```

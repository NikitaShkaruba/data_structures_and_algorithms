# Difference array

Difference arrays are used when you have a problem involving intervals of events.
You can think of the events all occuring along a number line.
This number line can represent time or position.
Events happen on this number line, with a starting and ending point.

The input will either provide a 2D array, where each inner array is in the format `[left, right, value]`, or some equivalent form.
The story behind the problem will usually be something along the lines of "between left and right, there is value of something".

Difference array pattern should create a difference array by adding `value` when `left` happens, and subtracting `value` when `right` happens.
Then we sort this difference array and iterate on it, executing

### [Example: 1094. Car Pooling](https://leetcode.com/problems/car-pooling/) 

> A car has room for `capacity` passengers, and is given an array `trips`. 
> Each trip is represented by `[numPassengers, from, to]`, which indicates that at `from`, it picks up `numPassengers`, then drops them off at `to`.
> Can it complete all the trips without holding more passengers than capacity at any time?

```golang
func carPooling(trips [][]int, capacity int) bool {
    // Create difference array
    events := make([]Event, 0)
    for _, t := range trips {
      events = append(events, Event{coordinate: t[1], capacityDiff: -t[0]})
      events = append(events, Event{coordinate: t[2], capacityDiff: t[0]})
    }

    // Sort it to make later iterations happen in a proper timeline
    sort.Slice(events, func(i, j int) bool {
      return events[i].coordinate < events[j].coordinate
    })

    i := 0
    for i < len(events) {
      // process the first event
      capacity += events[i].capacityDiff
      i++

      // process events on the same coordinate
      for i < len(events) && events[i-1].coordinate == events[i].coordinate {
        capacity += events[i].capacityDiff
        i++
      }

      if capacity < 0 {
        return false
      }
    }

    return true
}
```

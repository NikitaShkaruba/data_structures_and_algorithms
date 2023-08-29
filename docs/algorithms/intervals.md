# Intervals

Another common pattern is problems that involve intervals [start, end].
Usually, these problems can be solved by first sorting the input, then iterating over it while comparing adjacent intervals.

> [Example: 252. Meeting Rooms](https://leetcode.com/problems/meeting-rooms/)
> Given an array of meeting times intervals where intervals[i] = [start, end] indicates the i'th meeting runs from [start, end), determine if one person could attend all meetings.
> For example, given intervals = [[0, 30], [5, 10], [15, 20]], return false. If you attend the [0, 30] meeting, then you cannot attend the other two.

The brute force approach would be to look at every pair of meetings and check if there is overlap, which would take `O(n^2)`, where `n` is the number of meetings.
We can improve on this by observing that if there is a conflict, then the conflicting meetings would be adjacent if we sorted the input.
If we sort the meetings by their start time, then the meetings are in the order in which they should be attended.
If the i'th meeting starts before the (i-1)'th meeting ends, then there is a conflict, because you would need to leave (i-1)'th meeting early to be on time for the i'th meeting.
The main take of this approach: if you want to  find if intervals intersect, sort them, and iterate from left to right.

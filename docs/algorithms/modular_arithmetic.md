# Modular arithmetic

Sometimes, a problem will ask you to return an integer `answer mod x`. If you're on LeetCode, x is usually `10^9+7 = 1,000,000,007`.
The reason for this is that the real answer could be massive, too large to store in a variable.
In these problems, one thing you could do is calculate the actual massive answer, and then return `ans % MOD` at the end. But it can lead to overflows.

The interesting thing about the modulo operator is that you can use it at every step and still get the same answer.   
- With the division at the end `(11 * 15 * 26 * 43 * 62) % 7 = 11437140 % 7 = 1`
- With the division everywhere
  ```
  11 * 15 = 165 % 7 = 4
  4 * 26 = 104 % 7 = 6
  6 * 43 = 258 % 7 = 6
  6 * 62 = 372 % 7 = 1 // The same result
  ```

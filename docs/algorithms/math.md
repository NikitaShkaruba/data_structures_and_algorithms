# Math

Ahh, math problems. I like them the least! :D

### Greatest common divisor (GCD)
GCD helps to find number `c` that both input numbers `a` and `b` can be divided. Takes `O(logn)` time if you use [Euclid's algorithm](https://en.wikipedia.org/wiki/Euclidean_algorithm).
You can also get GCD using [primes factorization](#primes-factorization) bu getting primes for the first and the second number, and finding the greatest one.

### Primes factorization

Prime number only has two factors - one and the number itself.
Opposite of prime is composite.
We can get prime factors of number with the [Sieve of Eratosthenes](https://www.geeksforgeeks.org/sieve-of-eratosthenes/) algorithm which takes `O(n*log(log(n)))` to build.
Afterwards, `GetPrimeFactors` works in `O(logn)` time.

Source code: [src/algorithms/math.go](../../src/algorithms/math.go)

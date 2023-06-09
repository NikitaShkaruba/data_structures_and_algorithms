package algorithms

////////////////////////// Greatest common divider (GCD) //////////////////////////
// You can find this algorithms overview in docs/algorithms/math.md

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

////////////////////////// Primes factorization //////////////////////////
// You can find this algorithms overview in docs/algorithms/math.md

func GetPrimeFactors(number int, maxNumber int) []int {
	sieve := buildPrimeDividers(maxNumber)
	primeFactors := make([]int, 0)

	for number != 1 {
		factor := sieve[number]
		for number%factor == 0 {
			number /= factor
		}

		primeFactors = append(primeFactors, factor)
	}

	return primeFactors
}

func buildPrimeDividers(maxNumber int) []int {
	sieve := make([]int, maxNumber+1)

	for i := 2; i <= maxNumber; i++ {
		if sieve[i] != 0 {
			continue
		}

		sieve[i] = i
		for j := i * i; j <= maxNumber; j += i {
			if sieve[j] == 0 {
				sieve[j] = i
			}
		}
	}

	return sieve
}

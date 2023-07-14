package algorithms

////////////////////////// Prime factors //////////////////////////
// You can find this data structure overview in docs/algorithms/math.md

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

package stdlib

func GetIntegerLength(num int) int {
	if num == 0 {
		return 1
	}

	num = Abs(num)

	length := 0
	for num != 0 {
		num /= 10
		length++
	}

	return length
}

func GetIntegerDigit(num, position int) int {
	divider := Pow(10, position)
	return Abs(num) / divider % 10
}

package calaculator_go

func Sum(number1, number2 int) int {
	return number1 + number2
}

func Kali(number1, number2 int) int {
	return number1 * number2
}

func Kurang(number1, number2 int) int {
	return number1 - number2
}

func Bagi(number1, number2 int) int {
	if number2 == 0 {
		return 0
	}
	return number1 / number2
}

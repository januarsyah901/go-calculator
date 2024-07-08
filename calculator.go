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
func Modulo(number1, number2 int) int {
	if number2 == 0 {
		return 0
	}
	return number1 % number2
}

func Pangkat(number1, number2 int) int {
	result := 1
	for i := 0; i < number2; i++ {
		result *= number1
	}
	return result
}

package mathutil

import "errors"

var ErrDivByZero = errors.New("деление на ноль")

// Divide делит a на b и возвращает ошибку при b == 0
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivByZero
	}
	return a / b, nil
}

// Sum складывает произвольное количество целых чисел
func Sum(nums ...int) int {
	acc := 0
	for _, n := range nums {
		acc += n
	}
	return acc
}

// Multiply перемножает два целых числа (Обязательное задание №2)
func Multiply(a, b int) int {
	return a * b
}

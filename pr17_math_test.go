package mathutil

import (
	"fmt"
	"math"
	"testing"
)

// 1. Табличный тест для Divide (Обязательное задание №1)
func TestDivideTable(t *testing.T) {
	cases := []struct {
		title   string
		a, b    float64
		want    float64
		wantErr bool
	}{
		{"Целочисленное деление", 10, 2, 5.0, false},
		{"Отрицательные числа", -15, 3, -5.0, false},
		{"Оба отрицательные", -12, -4, 3.0, false},
		{"Дробный результат", 7, 2, 3.5, false},
		{"Периодическая дробь", 1, 3, 0.333333, false},
		{"Деление на ноль", 10, 0, 0, true},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			got, err := Divide(c.a, c.b)
			if (err != nil) != c.wantErr {
				t.Fatalf("Divide(%v, %v) ошибка = %v, ожидалась ошибка = %v", c.a, c.b, err, c.wantErr)
			}
			if !c.wantErr && math.Abs(got-c.want) > 1e-5 {
				t.Errorf("Divide(%v, %v) = %f, ожидалось %f", c.a, c.b, got, c.want)
			}
		})
	}
}

// 2. Табличный тест для Sum
func TestSumTable(t *testing.T) {
	cases := []struct {
		title string
		in    []int
		want  int
	}{
		{"Положительные", []int{10, 20, 30}, 60},
		{"С отрицательными", []int{-5, 10, -5}, 0},
		{"Пустой список", []int{}, 0},
		{"Одно число", []int{99}, 99},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			if got := Sum(c.in...); got != c.want {
				t.Errorf("Sum(%v) = %d, ожидалось %d", c.in, got, c.want)
			}
		})
	}
}

// 3. Тест для Multiply
func TestMultiply(t *testing.T) {
	if got := Multiply(-6, 7); got != -42 {
		t.Errorf("Multiply(-6, 7) = %d, ожидалось -42", got)
	}
}

// 4. Бенчмарки производительности (Обязательное задание №2)
func BenchmarkSum(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		_ = Sum(nums...)
	}
}

func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Multiply(12345, 67890)
	}
}

// 5. Документирующий пример-тест (Дополнительное задание №3)
func ExampleSum() {
	fmt.Println(Sum(1, 2, 3, 4))
	fmt.Println(Sum(-10, 10))
	// Output:
	// 10
	// 0
}

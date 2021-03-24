package fibonacci

// Метод минимизации - Фибоначчи
import (
	"fmt"
	"math"
)

// Function in task
func F(x float64) float64 {
	return 3*x - 5*math.Log(x)
}

// Генерация чисел Фибоначчи
func makeFibonacci(n int) []float64 {
	f := make([]float64, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f
}

func main() {
	var a float64 = 1 // Начало отрезка
	var b float64 = 3 // Конец отрезка
	var Eps float64 = 0.0001
	var d float64 = (Eps / 2) * 0.8 // Константа различимости

	var x1 float64 // Первая точка для сравнения
	var x2 float64 // Вторая точка для сравнения
	var F1 float64 // Значение функции от первой точки
	var F2 float64 // Значение функции от второй точки

	var Fib []float64 = makeFibonacci(30) // Числа Фибоначчи
	var Fn float64 = (b - a) / Eps
	var n int = 0 // Количество чисел Фибоначчи
	var k int = 1 // Количество итераций

	// Находим количество чисел Фибоначчи n
	for i := 1; i < len(Fib); i++ {
		if Fib[i] >= Fn {
			n = i
			Fn = Fib[i]
			break
		}
	}

	// Выполняем, пока не дойдем до пред-пред-последнего числа
	for k != n-2 {
		x1 = a + Fib[n-2]/Fn*(b-a)
		x2 = a + Fib[n-1]/Fn*(b-a)
		F1 = F(x1)
		F2 = F(x2)

		if F1 < F2 {

			// "Сдвигаемся" в сторону лучшего участка
			b = x2
			x2 = x1
			F2 = F1

			// Берем новое значение
			x1 = a + (Fib[n-k-1]/Fib[n-k])*(b-a)

			if k == n-2 {
				break
			} else {
				F1 = F(x1)
			}

		} else if F1 > F2 {

			// "Сдвигаемся" в сторону лучшего участка
			a = x1
			x1 = x2
			F1 = F2

			// Берем новое значение
			x2 = a + (Fib[n-k-2]/Fib[n-k])*(b-a)

			if k == n-2 {
				break
			} else {
				F2 = F(x2)
			}
		}
		k++
	}

	x2 = x1 + d
	if F(x1) < F(x2) {
		b = x2
	} else if F(x1) > F(x2) {
		a = x1
	}

	var xmin float64 = 0 // Итоговая точка минимума
	var Fmin float64 = 0 // Итоговый минимум функции

	// Проверяем, какая из границ дает меньшее значение функции и берем ее
	if F(a) < F(b) {
		Fmin = F(a)
		xmin = a
	} else if F(a) > F(b) {
		Fmin = F(b)
		xmin = b
	}

	fmt.Println("Минимум функции: ", Fmin)
	fmt.Println("Точка минимума: ", xmin)
	fmt.Println("Количество итераций: ", k)
}

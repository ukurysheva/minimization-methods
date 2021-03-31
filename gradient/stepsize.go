package main

// Градиентный минимизации - с постоянным шагом
import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/diff/fd"
)

// Function in task
func Fn(x []float64) float64 {
	return math.Pow((x[0]-5), 2)*math.Pow((x[1]-4), 2) + math.Pow((x[0]-5), 2) + math.Pow((x[1]-4), 2) + 1
}

// Function find Mod of two x
func Mod(x []float64) float64 {
	var sum = math.Pow(x[0], 2) + math.Pow(x[1], 2)
	return math.Cbrt(sum)
}

func findMin(M float64, ylocal float64, Eps1 float64, Eps2 float64) (float64, []float64, float64, float64) {

	var x01 float64 = 7      // Координата начальной точки
	var x02 float64 = 7      // Координата начальной точки
	x := []float64{x01, x02} // Начальная точка
	var xlocal []float64 = x
	var k float64 = 0  // Счетчик итераций
	var kf float64 = 0 // Счетчик вычисления функций
	x2 := []float64{}
	dst := make([]float64, 2) // Массив, в который будет помещен результат

	// Вычисляем градиент в начальной точке
	res := fd.Gradient(dst, Fn, xlocal, &fd.Settings{
		Step: 1e-3,
	})

	for Mod(res) >= Eps1 {
		if k < M {
			res = fd.Gradient(dst, Fn, xlocal, &fd.Settings{
				Step: 1e-3,
			})

			var yIsFine bool = false
			for !yIsFine {
				var x11 float64 = xlocal[0] - ylocal*res[0]
				var x12 float64 = xlocal[1] - ylocal*res[1]
				x2 = []float64{x11, x12} // Новая точка
				if Fn(x2)-Fn(xlocal) < 0 {
					yIsFine = true
				} else {
					ylocal = ylocal / 2
				}
				// Увеличиваем счетчик вычисления функции
				kf += 2
			}

			var x1diff float64 = x2[0] - xlocal[0]
			var x2diff float64 = x2[1] - xlocal[1]
			xdiff := []float64{x1diff, x2diff} // Начальная точка
			if Mod(xdiff) <= Eps2 && math.Abs(Fn(x2)-Fn(xlocal)) <= Eps2 {
				xlocal = x2
				break
			} else {
				k++
				xlocal = x2
				res = fd.Gradient(dst, Fn, xlocal, &fd.Settings{
					Step: 1e-3,
				})
			}
		} else {
			break
		}

	}

	return Fn(xlocal), xlocal, k, kf
}
func main() {

	var M float64 = 100 // Максимальное число итераций

	var y float64 = 1 // Лямбда
	var Eps1 float64 = 0.1

	Fmin, x, k, kf := findMin(M, y, Eps1, Eps1)
	fmt.Println("EPS: ", Eps1)
	fmt.Println("Минимум функции: ", Fmin)
	fmt.Println("Точка минимума: ", x)
	fmt.Println("Количество шагов: ", k)
	fmt.Println("Количество вычислений функции: ", kf, "\n")
}

package main

// Градиентный минимизации наискорейшего спуска
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

func ysearch(x []float64, grad []float64) float64 {
	var a float64 = 0 // Начало отрезка
	var b float64 = 1 // Конец отрезка
	var Eps float64 = 0.0001
	var N float64 = (b - a) / Eps // Количество отрезков
	var Fmin float64 = Fn(x)      // Минимум функции
	var ymin float64 = 0          // Гамма минимальная
	var i float64 = (b - a) / N   // Расстояние между точками
	var x11 float64
	var x12 float64
	for y := a + i; y <= b; y += i {
		x11 = x[0] - y*grad[0]
		x12 = x[1] - y*grad[1]
		x2 := []float64{x11, x12} // Начальная точка
		if Fn(x2) <= Fmin {
			Fmin = Fn(x2)
			ymin = y
		}
	}

	return ymin
}
func findMin(M float64, y float64, Eps1 float64, Eps2 float64) (float64, []float64, float64, float64) {

	var x01 float64 = 7      // Координата начальной точки
	var x02 float64 = 7      // Координата начальной точки
	x := []float64{x01, x02} // Начальная точка
	var k float64 = 0        // Счетчик
	var kf float64 = 0       // Счетчик вычисления функций
	x2 := []float64{}

	dst := make([]float64, 2) // Массив, в который будет помещен результат

	// Вычисляем градиент в начальной точке
	res := fd.Gradient(dst, Fn, x, &fd.Settings{
		Step: 1e-3,
	})

	for Mod(res) >= Eps1 {
		if k < M {
			res = fd.Gradient(dst, Fn, x, &fd.Settings{
				Step: 1e-3,
			})

			y = ysearch(x, res)
			var x11 float64 = x[0] - y*res[0]
			var x12 float64 = x[1] - y*res[1]
			x2 = []float64{x11, x12} // Начальная точка

			var x1diff float64 = x2[0] - x[0]
			var x2diff float64 = x2[1] - x[1]
			xdiff := []float64{x1diff, x2diff} // Начальная точка
			kf += 2
			if Mod(xdiff) <= Eps2 && math.Abs(Fn(x2)-Fn(x)) <= Eps2 {
				x = x2
				break
			} else {
				k++
				x = x2
				res = fd.Gradient(dst, Fn, x, &fd.Settings{
					Step: 1e-3,
				})

			}
		} else {
			break
		}

	}

	return Fn(x), x, k, kf
}
func main() {

	var M float64 = 100 // Чисто итераций
	var y float64 = 1   // Гамма
	var Eps1 float64 = 0.1
	Fmin, x, k, kf := findMin(M, y, Eps1, Eps1)
	fmt.Println("EPS: ", Eps1)
	fmt.Println("Минимум функции: ", Fmin)
	fmt.Println("Точка минимума: ", x)
	fmt.Println("Количество шагов: ", k)
	fmt.Println("Количество вычислений функции: ", kf, "\n")
}

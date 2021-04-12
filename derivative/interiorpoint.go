package interiorpoint

// Метод минимизации с исп.производной - метод средней точки
import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/diff/fd"
)

// Function in task
func F(x float64) float64 {
	return (math.Pow(x, 2) - x*(math.Pow(math.E, -x)))
}


func main() {
	var a float64 = 0 // Начало отрезка
	var b float64 = 1 // Конец отрезка
	var Eps float64 = 0.0005

	var k int = 0 
	var w int = 0 // Количество вычислений производной
	var xi float64 = (a + b) / 2 // Вычисляемая точка в данный момент

	var Fd float64 = fd.Derivative(F, xi, &fd.Settings{
		Formula: fd.Forward,
		Step:    1e-3,
	})

    w++
	 
	for math.Abs(Fd) > Eps {

		if Fd > 0 {
			k++
			b = xi
		} else {
			k++
			a = xi
		}

		xi = (a + b) / 2
		Fd = fd.Derivative(F, xi, &fd.Settings{
			Formula: fd.Forward,
			Step:    1e-3,
		})
        w++
	}

	var Fmin float64 = F(xi)
	fmt.Println("Минимум функции: ", Fmin)
	fmt.Println("Точка минимума: ", xi)
	fmt.Println("Количество вычислений Fdx: ", w)
}

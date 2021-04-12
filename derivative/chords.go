package chords

// Метод минимизации с исп.производной - метод хорд
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
	var w int = 0    // Количество вычислений производной

	var Fda float64 = fd.Derivative(F, a, &fd.Settings{
		Formula: fd.Forward,
		Step:    1e-3,
	})
    w++

	var Fdb float64 = fd.Derivative(F, b, &fd.Settings{
		Formula: fd.Forward,
		Step:    1e-3,
	})
    w++

	var xi float64 = a - (Fda/(Fda - Fdb))*(a-b) // Вычисляемая точка в данный момент

	var Fdx float64 = fd.Derivative(F, xi, &fd.Settings{
		Formula: fd.Forward,
		Step:    1e-3,
	})
	w++

	if (Fda * Fdb < 0) { 

		for math.Abs(Fdx)>Eps {

			if Fdx > 0 {
				k++
				b = xi
				Fdb = Fdx
			} else {
				k++
				a = xi
				Fda = Fdx
			}

			xi = a - (Fda/(Fda - Fdb))*(a-b)

			Fdx = fd.Derivative(F, xi, &fd.Settings{
				Formula: fd.Forward,
				Step:    1e-3,
			})
			w++
		}
	}

	if Fda > 0 && Fdb > 0 {
		xi = a
	} else if Fda < 0 && Fdb < 0 {
		xi = b
	} else if Fda*Fdb == 0 {
		if Fda == 0 {
			xi = a
		} else if Fdb == 0 {
			xi = b
		}
	}

    var Fmin float64 = F(xi)
	
	fmt.Println("Минимум функции: ", Fmin)
	fmt.Println("Точка минимума: ", xi)
	fmt.Println("Количество вычислений Fdx: ", w)
}

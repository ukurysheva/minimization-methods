package main

// Descent gradient minimization method
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

// Function finds the best y
func ysearch(x []float64, grad []float64) float64 {
	var a float64 = 0 // Begin point
	var b float64 = 1 // End point
	var Eps float64 = 0.0001
	var N float64 = (b - a) / Eps
	var Fmin float64 = Fn(x)    // Min value of function
	var ymin float64 = 0        // Y min
	var i float64 = (b - a) / N // Step
	var x11 float64
	var x12 float64
	for y := a + i; y <= b; y += i {
		x11 = x[0] - y*grad[0]
		x12 = x[1] - y*grad[1]
		x2 := []float64{x11, x12}
		if Fn(x2) <= Fmin {
			Fmin = Fn(x2)
			ymin = y
		}
	}

	return ymin
}

// Function finding min value
func findMin(M float64, y float64, Eps1 float64, Eps2 float64) (float64, []float64, float64, float64) {

	var x01 float64 = 7 // Begin point - x
	var x02 float64 = 7 // Begin point - y
	x := []float64{x01, x02}
	var k float64 = 0  // Iterator
	var kf float64 = 0 // Iterator of function calculation
	x2 := []float64{}

	dst := make([]float64, 2) // The result array

	// Find gradient in the beginning point
	res := fd.Gradient(dst, Fn, x, &fd.Settings{
		Step: 1e-3,
	})

	for Mod(res) >= Eps1 {
		if k < M {
			res = fd.Gradient(dst, Fn, x, &fd.Settings{
				Step: 1e-3,
			})

			y = ysearch(x, res)

			// New point x
			var x11 float64 = x[0] - y*res[0]
			var x12 float64 = x[1] - y*res[1]
			x2 = []float64{x11, x12}

			var x1diff float64 = x2[0] - x[0]
			var x2diff float64 = x2[1] - x[1]
			xdiff := []float64{x1diff, x2diff} // Find norm
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

	var M float64 = 100 // The maximum count if iterations
	var y float64 = 1
	var Eps1 float64 = 0.1
	Fmin, x, k, kf := findMin(M, y, Eps1, Eps1)
	fmt.Println("EPS: ", Eps1)
	fmt.Println("Min value of F: ", Fmin)
	fmt.Println("Min point: ", x)
	fmt.Println("Iterations: ", k)
	fmt.Println("Iterations of F calculations: ", kf, "\n")
}

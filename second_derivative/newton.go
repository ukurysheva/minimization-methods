package main

// Newton minimization method of function with 2 params
import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/diff/fd"
	"gonum.org/v1/gonum/mat"
)

// Function in task
func Fn(x []float64) float64 {
	return 100*math.Pow(x[0]-math.Pow(x[1], 2), 2) + math.Pow(1-x[1], 2)
}

// Function find Mod of x vector
func Mod(x []float64) float64 {
	var sum = math.Pow(x[0], 2) + math.Pow(x[1], 2)
	return math.Cbrt(sum)
}

// Function check if matrix positive definite
func checkIfMatrPos(m mat.Dense) bool {
	var isPos = false
	// Find the D1 and D2 - minors of matrix
	var D1 = m.At(0, 0)
	var D2 = m.At(0, 0)*m.At(1, 1) - m.At(0, 1)*m.At(1, 0)
	if D1 > 0 && D2 > 0 {
		isPos = true
	}
	return isPos
}

// Function mupliply array[2] and matrix 2x2
func multiplyMatrix(x []float64, m mat.Dense) []float64 {
	dim := make([]float64, 2) // The result array
	dim[0] = m.At(0, 0)*x[0] + m.At(1, 0)*x[1]
	dim[1] = m.At(0, 1)*x[0] + m.At(1, 1)*x[1]

	return dim
}

// Function finding min value
func findMin(M float64, Eps1 float64, Eps2 float64) (float64, []float64, float64, float64) {

	var x01 float64 = 2 // Begin point - x
	var x02 float64 = 3 // Begin point - y
	x := []float64{x01, x02}
	var k float64 = 0  // Iterator of steps
	var kf float64 = 0 // Iterator of function calculation

	fmt.Println("-----------------")
	fmt.Println("BEGIN POINT: ", x)

	grad := make([]float64, 2)
	// Find gradient in the beginning point
	fd.Gradient(grad, Fn, x, &fd.Settings{
		Step: 1e-3,
	})

	for Mod(grad) >= Eps1 {
		if k < M {

			grad := make([]float64, 2)
			// Find gradient in the beginning point
			fd.Gradient(grad, Fn, x, &fd.Settings{
				Step: 1e-3,
			})
			// Find Hessian matrix
			var Hes = mat.SymDense{}
			fd.Hessian(&Hes, Fn, x, &fd.Settings{
				Step: 1e-3,
			})
			// Convert SymDense to Matrix type
			a := mat.NewDense(2, 2, []float64{
				Hes.At(0, 0), Hes.At(1, 0),
				Hes.At(0, 1), Hes.At(1, 1),
			})

			// Make an inverse matrix of the Hessian matrix
			var HesInv mat.Dense
			err := HesInv.Inverse(a)
			if err != nil {
				fmt.Printf("A is not invertible: %v", err)
				break
			}

			var t float64 = 0           // used to compensate negative params
			d := make([]float64, 2)     // the diff between xk and xk+1
			tdiff := make([]float64, 2) // the diff between xk and xk+1
			xnext := make([]float64, 2) // xk+1
			xnext[0] = 0
			xnext[1] = 0

			// Check if teterminant is positive - multiply gradient with Hessian
			if checkIfMatrPos(HesInv) {
				d = multiplyMatrix(grad, HesInv)
				d[0] = -d[0]
				d[1] = -d[1]
				t = 1
				xnext[0] = x[0] + t*d[0]
				xnext[1] = x[1] + t*d[1]
			} else {
				d = grad
				d[0] = -d[0]
				d[1] = -d[1]
				tdiff = multiplyMatrix(d, HesInv)
				xnext[0] = x[0] + tdiff[0]*d[0]
				xnext[1] = x[1] + tdiff[1]*d[1]
			}

			// Find norm vector - diff between xk and xk+1
			var x1diff float64 = xnext[0] - x[0]
			var x2diff float64 = xnext[1] - x[1]
			xdiff := []float64{x1diff, x2diff}

			kf += 2
			if Mod(xdiff) <= Eps2 && math.Abs(Fn(xnext)-Fn(x)) <= Eps2 {
				x = xnext
				break
			} else {
				x = xnext
				k++
			}
		} else {
			break
		}
	}
	return Fn(x), x, k, kf
}
func main() {

	var M float64 = 100    // The maximum count of iterations
	var Eps1 float64 = 0.1 // Approximate value

	Fmin, x, k, kf := findMin(M, Eps1, Eps1)

	fmt.Println("EPS: ", Eps1)
	fmt.Println("Min value of F: ", Fmin)
	fmt.Println("Min point: ", x)
	fmt.Println("Iterations: ", k)
	fmt.Println("Iterations of F calculations: ", kf)
}

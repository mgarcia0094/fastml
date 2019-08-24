package fastml

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func CreateVector() {
	u := mat.NewVecDense(3, []float64{1, 2, 3})
	println("u: ")
	matPrint(u)
}


package fastml

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func matPrint(X mat.Matrix) {
 fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
 fmt.Printf("%v\n", fa)
}

func CreateVector() {
	u := mat.NewVecDense(3, []float64{1, 2, 3})
	fmt.Printf("u: ")
	matPrint(u)
}



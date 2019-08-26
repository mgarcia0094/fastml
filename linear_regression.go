package fastml

import (
	"errors"
	"gonum.org/v1/gonum/mat"
)


//Linear Regression struct
type linearRegression struct {
	l_r float64
	n_iter uint
	epsilon float64
	gamma *mat.VecDense
	beta *mat.VecDense
}

//Constructor linear regression
//
//Params:
//	l_r: learning rate
//	n_iter: number of iteration fit
//	epsilon : error max tolerate
func NewLinearRegression(l_r float64, n_iter uint, epsilon float64) *linearRegression {
	return &linearRegression{l_r, n_iter, epsilon, nil, nil}
}

func (l_r *linearRegression) Fit (X_train *mat.VecDense, Y_train *mat.VecDense) error {
	return errors.New("FastML error sample")
}

func (l_r *linearRegression) Score (X_test *mat.Matrix, Y_test *mat.VecDense) (error, float64) {
	return nil, -1
}

func (l_r *linearRegression) Predict (X_predict *mat.VecDense) (error, *mat.VecDense) {
	return nil, X_predict
}



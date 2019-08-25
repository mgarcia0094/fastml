package fastml

import (
	"errors"
	"gonum.org/v1/gonum/mat"
)

type linearRegression struct {
	coef_ float64
	intercept_ float64
}

func NewLinearRegression() (*linearRegression) {
	return &linearRegression{0, 0}
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



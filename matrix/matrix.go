// https://gist.github.com/n1try/c5082f0f1db7f4abcb6d995dc275fe7f
/* Simple Matrix multiplication and transpostion with Go */
package matrix

import (
	"errors"
)

type Matrix [][]float64

func Transpose(x Matrix) Matrix {
	out := make(Matrix, len(x[0]))
	for i := 0; i < len(x); i ++ {
		for j := 0; j < len(x[0]); j ++ {
			out[j] = append(out[j], x[i][j])
		}
	}
	return out
}

func Dot(x, y Matrix) (Matrix) {
	if len(x[0]) != len(y) {
    panic(errors.New("Can't do matrix multiplication."))
	}

	out := make(Matrix, len(x))
	for i := 0; i < len(x); i ++ {
		for j := 0; j < len(y); j ++ {
			if len(out[i]) < 1 {
				out[i] = make([]float64, len(y))
			}
			out[i][j] += x[i][j] * y[j][i]
		}
	}
	return out
}

// https://gist.github.com/n1try/c5082f0f1db7f4abcb6d995dc275fe7f
package matrix

import (
  "fmt"
	"errors"
)

var (
  dimerr = errors.New("Dimension mismatch")
)

type Matrix [][]float64

func NewMatrix(r, c int, x []float64) Matrix{
  if r < 0 || c < 0{
    panic("Negative dimension matrix")
  }
  if r*c != len(x) && x != nil{
    panic(dimerr)
  }
  if x == nil{
    x = make([]float64, r*c)
  }
  out := make(Matrix, r)
  for i := range out{
    out[i] = x[c*i:(c*i)+c]
  }
  return out
}

func Transpose(x Matrix) Matrix {
	out := make(Matrix, len(x[0]))
	for i := 0; i < len(x); i ++ {
		for j := 0; j < len(x[0]); j ++ {
			out[j] = append(out[j], x[i][j])
		}
	}
	return out
}

func Dot(a, b Matrix) Matrix{
  // https://github.com/shiffman/Neural-Network-p5/blob/master/matrix.js
  out := NewMatrix(len(a), len(b[0]), nil)
  for i := 0; i < len(a); i++{
    for j := 0; j < len(b[0]); j++{
      sum := 0.0
      for k := 0; k < len(a[0]); k++{
        sum += a[i][k] * b[k][j]
      }
      out[i][j] = sum
    }
  }
	return out
}

func Add(m ...Matrix) Matrix {
  for i := 1; i<len(m); i++{
    if len(m[i-1]) != len(m[i]) || len(m[i-1][0]) != len(m[i][0]){
      panic(dimerr)
    }
  }
  out := NewMatrix(len(m[0]), len(m[0][0]), nil)
  for _, v := range m{
    for i, w := range v{
      for j := range w{
        out[i][j] += v[i][j]
      }
    }
  }
  return out
}

func AddVal(m Matrix, f float64) Matrix{
  out := NewMatrix(m.Rows(), m.Cols(), nil)
  for i := 0; i < m.Rows(); i++{
    for j := 0; j < m.Cols(); j++{
      out[i][j] = m[i][j] + f
    }
  }
  return out
}

func Sub(m ...Matrix) Matrix {
  for i := 1; i<len(m); i++{
    if len(m[i-1]) != len(m[i]) || len(m[i-1][0]) != len(m[i][0]){
      panic(dimerr)
    }
  }
  out := Copy(m[0])
  for k := 1; k < len(m); k++{
    v := m[k]
    for i := 0; i < v.Rows(); i++{
      for j := 0; j < v.Cols(); j++{
        out[i][j] -= v[i][j]
      }
    }
  }
  return out
}

func ValSub(f float64, m Matrix) Matrix{
  out := NewMatrix(m.Rows(), m.Cols(), nil)
  for i := 0; i < m.Rows(); i++{
    for j := 0; j < m.Cols(); j++{
      out[i][j] = f - m[i][j]
    }
  }
  return out
}

func SubVal(m Matrix, f float64) Matrix{
  out := NewMatrix(m.Rows(), m.Cols(), nil)
  for i := 0; i < m.Rows(); i++{
    for j := 0; j < m.Cols(); j++{
      out[i][j] = m[i][j] - f
    }
  }
  return out
}

func Mul(m ...Matrix) Matrix {
  for i := 1; i<len(m); i++{
    if len(m[i-1]) != len(m[i]) || len(m[i-1][0]) != len(m[i][0]){
      panic(dimerr)
    }
  }
  out := NewMatrix(len(m[0]), len(m[0][0]), nil)
  for i, v := range out{
    for j := range v{
      out[i][j] = 1.0
    }
  }
  for _, v := range m{
    for i, w := range v{
      for j := range w{
        out[i][j] *= v[i][j]
      }
    }
  }
  return out
}

func MulVal(m Matrix, f float64) Matrix{
  out := NewMatrix(m.Rows(), m.Cols(), nil)
  for i := 0; i < m.Rows(); i++{
    for j := 0; j < m.Cols(); j++{
      out[i][j] = m[i][j] * f
    }
  }
  return out
}

func (m Matrix) Rows() int {
  return len(m)
}

func (m Matrix) Cols() int {
  return len(m[0])
}

func Copy(m Matrix) Matrix {
  out := NewMatrix(m.Rows(), m.Cols(), nil)
  for i := 0; i < m.Rows(); i++{
    for j := 0; j < m.Cols(); j++{
      out[i][j] = m[i][j]
    }
  }
  return out
}

func (m Matrix) String() string {
  out := ""
  for i := 0; i < m.Rows()-1; i++{
    out += fmt.Sprintf("%v\n", m[i])
  }
  return fmt.Sprintf("[%v%v]", out, m[m.Rows()-1])
}

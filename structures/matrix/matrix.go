package matrix

import "fmt"

type Matrix [][]float64

func New(rows, cols int) Matrix {
	matrix := make(Matrix, rows)
	for i := range matrix {
		matrix[i] = make([]float64, cols)
	}

	return matrix
}

func (m Matrix) Fill(v float64) Matrix {
	for i, rows := range m {
		for j := range rows {
			m[i][j] = v
		}
	}

	return m
}

func (m Matrix) Max() (max float64) {
	for _, rows := range m {
		for _, x := range rows {
			if x > max {
				max = x
			}
		}
	}

	return
}

// String to prepare matrix to output.
func (m Matrix) String() (s string) {
	for _, rows := range m {
		for _, x := range rows {
			s += fmt.Sprintf("%v ", x)
		}

		s += "\n"
	}

	return
}

func (m Matrix) AddMatrix(m2 Matrix) Matrix {
	for i, rows := range m {
		for j := range rows {
			m[i][j] += m2[i][j]
		}
	}

	return m
}

func (m Matrix) MulMatrix(m2 Matrix) (m3 Matrix) {
	m3 = New(len(m), len(m[0]))

	for i, rows := range m {
		for j, x := range rows {
			for k := range m2 {
				m3[i][j] = x * m2[k][j]
			}
		}
	}

	return
}

func (m Matrix) MulScalar(scalar float64) Matrix {
	for i, rows := range m {
		for j := range rows {
			m[i][j] *= scalar
		}
	}

	return m
}

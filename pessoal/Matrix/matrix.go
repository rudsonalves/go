package main

import "fmt"

type matrix [][]float64

func setMatrix(m, n int, e ...float64) (x matrix) {
	if len(x) == 0 {
		x = setZeros(m, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			x[i][j] = e[i*n+j]
		}
	}
	return
}

func setZeros(m, n int) (x matrix) {
	for i := 0; i < m; i++ {
		l := []float64{}
		for j := 0; j < n; j++ {
			l = append(l, 0)
		}
		x = append(x, l)
	}
	return
}

func (x matrix) len() (m, n int) {
	m = len(x)
	n = len(x[0])
	return
}

func (x matrix) Add(y matrix) (r matrix) {
	m, n := x.len()
	my, ny := y.len()
	if m != my || n != ny {
		fmt.Println("Add error: x and y have different sizes")
		return
	}
	r = setZeros(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			r[i][j] = x[i][j] + y[i][j]
		}
	}
	return
}

func (x matrix) String() string {
	m, n := x.len()
	out := ""
	for i := 0; i < m; i++ {
		line := "["
		for j := 0; j < n; j++ {
			line += fmt.Sprintf("  %+4.2f ", x[i][j])
		}
		out += line + "]\n"
	}
	return out
}

func (x matrix) Mult(y matrix) (r matrix) {
	mx, nx := x.len()
	my, ny := y.len()
	if mx != ny || nx != my {
		fmt.Println("Mult error: x and y have different sizes")
		return
	}
	m, n := mx, ny
	r = setZeros(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < nx; k++ {
				r[i][j] += x[i][k] * y[k][j]
			}
		}
	}
	return
}

func main() {

	x1 := setMatrix(2, 3, 1, 2, 3, -3, -2, -1, -1, 3, -2)
	y1 := setMatrix(3, 3, 0, 0, 1, 2, 0, -1, 6, 5, 2)
	fmt.Println(x1.len())
	fmt.Println(x1)
	fmt.Println(y1.len())
	fmt.Println(y1)
	fmt.Println(x1.Add(x1))
	fmt.Println(x1.Mult(y1))
}

package main

import "errors"

func TraverseSpiral2D(m [][]int) ([]int, error) {
	dim, ok := measureDim(m)
	if ok != true {
		return nil, errors.New("wrong dimension")
	}
	fromX := 0
	fromY := 1
	toX := dim - 1
	toY := dim - 1
	cx := fromX
	cy := fromY - 1
	step := 1
	flat := make([]int, 0, dim*dim)

	for {
		r := make(chan int)
		go traverseAxis(r, fromX, toX, step)
		for cx = range r {
			flat = append(flat, m[cy][cx])
		}
		if fromX == toX {
			break
		}
		r = make(chan int)
		go traverseAxis(r, fromY, toY, step)
		for cy = range r {
			flat = append(flat, m[cy][cx])
		}
		fromX, toX = toX-step, fromX
		fromY, toY = toY-step, fromY

		step = step * -1
	}
	return flat, nil
}

func traverseAxis(r chan<- int, from, to, step int) {
	if step > 0 {
		for i := from; i <= to; i += step {
			r <- i
		}
		close(r)
		return
	}

	for i := from; i >= to; i += step {
		r <- i
	}
	close(r)
}

func measureDim(m [][]int) (dim int, ok bool) {
	h := len(m)
	if h == 0 {
		return 0, false
	}
	w := len(m[0])
	if w != h {
		return 0, false
	}
	for _, subm := range m {
		if len(subm) != w {
			return 0, false
		}
	}

	return h, true
}

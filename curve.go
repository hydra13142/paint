package paint

import "math"

// 绘制圆、椭圆
func (img *Image) Ellipse(smlX, smlY, bigX, bigY int) {
	if smlX == bigX || smlY == bigY {
		img.Line(smlX, smlY, bigX, bigY)
		return
	}
	if smlX > bigX {
		smlX, bigX = bigX, smlX
	}
	if smlY > bigY {
		smlY, bigY = bigY, smlY
	}

	rx := float64(bigX-smlX) / 2
	ry := float64(bigY-smlY) / 2

	w, h := (bigX-smlX)/2+1, (bigY-smlY)/2+1
	u := make([]bool, h)

	var a, b, c, d int
	var e, f float64
	if i := smlX + bigX; i > 0 {
		a, c = i/2, (i+1)/2
	} else {
		a, c = (i-1)/2, i/2
	}
	if j := smlY + bigY; j > 0 {
		b, d = j/2, (j+1)/2
	} else {
		b, d = (j-1)/2, j/2
	}
	e, f = 0.5, 0.5
	if a == c {
		e = 0
	}
	if b == d {
		f = 0
	}

	for i := 0; i < w; i++ {
		dx := float64(i) + e
		j := round(math.Sqrt(1-(dx*dx)/(rx*rx)) * ry)
		img.Pset(c+i, b-j)
		img.Pset(a-i, b-j)
		img.Pset(c+i, d+j)
		img.Pset(a-i, d+j)
		u[j] = true
	}
	for j := 0; j < h; j++ {
		if !u[j] {
			dy := float64(j) + f
			i := round(math.Sqrt(1-(dy*dy)/(ry*ry)) * rx)
			img.Pset(c+i, b-j)
			img.Pset(a-i, b-j)
			img.Pset(c+i, d+j)
			img.Pset(a-i, d+j)
		}
	}
}

// 绘制圆、椭圆的一部分，采用弧度表示弧线部分范围
func (img *Image) Arc(smlX, smlY, bigX, bigY int, frm, end float64) {
	if end-frm >= math.Pi*2 {
		img.Ellipse(smlX, smlY, bigX, bigY)
		return
	}
	if smlX > bigX {
		smlX, bigX = bigX, smlX
	}
	if smlY > bigY {
		smlY, bigY = bigY, smlY
	}
	for frm < 0 {
		frm += math.Pi * 2
	}
	for frm > math.Pi*2 {
		frm -= math.Pi * 2
	}
	for end < frm {
		end += math.Pi * 2
	}

	rx := float64(bigX-smlX) / 2
	ry := float64(bigY-smlY) / 2

	w, h := (bigX-smlX)/2+1, (bigY-smlY)/2+1
	u := make([]bool, h)

	var a, b, c, d int
	var e, f float64
	if i := smlX + bigX; i > 0 {
		a, c = i/2, (i+1)/2
	} else {
		a, c = (i-1)/2, i/2
	}
	if j := smlY + bigY; j > 0 {
		b, d = j/2, (j+1)/2
	} else {
		b, d = (j-1)/2, j/2
	}
	e, f = 0.5, 0.5
	if a == c {
		e = 0
	}
	if b == d {
		f = 0
	}

	var In func(float64) bool
	if end <= 2*math.Pi {
		In = func(ag float64) bool {
			return ag >= frm && ag <= end
		}
	} else {
		mdi := end - 2*math.Pi
		In = func(ag float64) bool {
			return ag >= frm || ag <= mdi
		}
	}

	for i := 0; i < w; i++ {
		dx := float64(i) + e
		dy := math.Sqrt(1-(dx*dx)/(rx*rx)) * ry
		ag := math.Atan2(dy, dx)
		j := round(dy)
		if In(ag) {
			img.Pset(c+i, b+j)
		}
		if In(math.Pi - ag) {
			img.Pset(a-i, b+j)
		}
		if In(math.Pi + ag) {
			img.Pset(a-i, d-j)
		}
		if In(math.Pi*2 - ag) {
			img.Pset(c+i, d-j)
		}
		u[j] = true
	}
	for j := 0; j < h; j++ {
		if !u[j] {
			dy := float64(j) + f
			dx := math.Sqrt(1-(dy*dy)/(ry*ry)) * rx
			ag := math.Atan2(dy, dx)
			i := round(dx)
			if In(ag) {
				img.Pset(c+i, b+j)
			}
			if In(math.Pi - ag) {
				img.Pset(a-i, b+j)
			}
			if In(math.Pi + ag) {
				img.Pset(a-i, d-j)
			}
			if In(math.Pi*2 - ag) {
				img.Pset(c+i, d-j)
			}
		}
	}
}

package paint

import (
	"image/color"
	"math"
)

func Round(x float64) int {
	if x >= 0 {
		if x-math.Floor(x) >= 0.5 {
			return int(x) + 1
		} else {
			return int(x)
		}
	} else {
		if x-math.Floor(x) >= 0.5 {
			return int(x)
		} else {
			return int(x) - 1
		}
	}
}

func Angle(x, y float64) float64 {
	if x > 0 {
		if y < 0 {
			return math.Atan(y/x) + math.Pi*2
		}
		return math.Atan(y / x)
	}
	if x < 0 {
		return math.Atan(y/x) + math.Pi
	}
	if y > 0 {
		return +math.Pi / 2
	}
	if y < 0 {
		return -math.Pi / 2
	}
	return 0
}

func Ellipse(img Image, l, t, r, b int, clr color.Color) {
	if l == r || t == b {
		Line(img, l, t, r, b, clr)
		return
	}
	if l > r {
		l, r = r, l
	}
	if t > b {
		t, b = b, t
	}

	rx := float64(r-l) / 2
	ry := float64(b-t) / 2

	w, h := (r-l)/2+1, (b-t)/2+1
	u := make([]bool, h)

	a, c := (l+r)/2, (l+r+1)/2
	b, d := (t+b)/2, (t+b+1)/2

	e, f := 0.5, 0.5
	if a == c {
		e = 0
	}
	if b == d {
		f = 0
	}

	for i := 0; i < w; i++ {
		dx := float64(i) + e
		j := Round(math.Sqrt(1-(dx*dx)/(rx*rx)) * ry)
		img.Set(c+i, b-j, clr)
		img.Set(a-i, b-j, clr)
		img.Set(c+i, d+j, clr)
		img.Set(a-i, d+j, clr)
		u[j] = true
	}
	for j := 0; j < h; j++ {
		if !u[j] {
			dy := float64(j) + f
			i := Round(math.Sqrt(1-(dy*dy)/(ry*ry)) * rx)
			img.Set(c+i, b-j, clr)
			img.Set(a-i, b-j, clr)
			img.Set(c+i, d+j, clr)
			img.Set(a-i, d+j, clr)
		}
	}
}

func Arc(img Image, l, t, r, b int, frm, end float64, clr color.Color) {
	if end-frm >= math.Pi*2 {
		Ellipse(img, l, t, r, b, clr)
		return
	}
	if l > r {
		l, r = r, l
	}
	if t > b {
		t, b = b, t
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

	rx := float64(r-l) / 2
	ry := float64(b-t) / 2

	w, h := (r-l)/2+1, (b-t)/2+1
	u := make([]bool, h)

	a, c := (l+r)/2, (l+r+1)/2
	b, d := (t+b)/2, (t+b+1)/2

	e, f := 0.5, 0.5
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
		ag := Angle(dx, dy)
		j := Round(dy)
		if In(ag) {
			img.Set(c+i, b+j, clr)
		}
		if In(math.Pi - ag) {
			img.Set(a-i, b+j, clr)
		}
		if In(math.Pi + ag) {
			img.Set(a-i, d-j, clr)
		}
		if In(math.Pi*2 - ag) {
			img.Set(c+i, d-j, clr)
		}
		u[j] = true
	}
	for j := 0; j < h; j++ {
		if !u[j] {
			dy := float64(j) + f
			dx := math.Sqrt(1-(dy*dy)/(ry*ry)) * rx
			ag := Angle(dx, dy)
			i := Round(dx)
			if In(ag) {
				img.Set(c+i, b+j, clr)
			}
			if In(math.Pi - ag) {
				img.Set(a-i, b+j, clr)
			}
			if In(math.Pi + ag) {
				img.Set(a-i, d-j, clr)
			}
			if In(math.Pi*2 - ag) {
				img.Set(c+i, d-j, clr)
			}
		}
	}
}

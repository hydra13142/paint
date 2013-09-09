package paint

import (
	"image/color"
	"math"
)

const PI = 3.1415926535897928264339

func round(x float64) int {
	if x-math.Ceil(x) >= 0.5 {
		return int(x) + 1
	} else {
		return int(x)
	}
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
	ox := float64(l+r) / 2
	oy := float64(t+b) / 2
	rx := float64(r-l) / 2
	ry := float64(b-t) / 2
	md := math.Sqrt(rx*rx + ry*ry)
	px := rx * rx / md
	py := ry * ry / md
	L, R := round(ox-px), round(ox+px)
	T, B := round(oy-py), round(oy+py)
	if (r-l)%2 == 0 {
		rx += 0.01
	}
	if (b-t)%2 == 0 {
		ry += 0.01
	}
	for i := L; i <= R; i++ {
		x := float64(i) - ox
		y := math.Sqrt(1-(x*x)/(rx*rx)) * ry
		img.Set(i, round(oy+y), clr)
		img.Set(i, round(oy-y), clr)
	}
	for j := T; j <= B; j++ {
		y := float64(j) - oy
		x := math.Sqrt(1-(y*y)/(ry*ry)) * rx
		img.Set(round(ox+x), j, clr)
		img.Set(round(ox-x), j, clr)
	}
}

func Arc(img Image, l, t, r, b int, frm, end float64, clr color.Color) {
	for end < frm {
		end += PI * 2
	}
	if end-frm >= PI*2 {
		Ellipse(img, l, t, r, b, clr)
		return
	}
	if l > r {
		l, r = r, l
	}
	if t > b {
		t, b = b, t
	}
	ox := float64(l+r) / 2
	oy := float64(t+b) / 2
	rx := float64(r-l) / 2
	ry := float64(b-t) / 2
	fx, fy := math.Cos(frm)*rx, math.Sin(frm)*ry
	if frm == end {
		img.Set(round(ox+fx), round(oy-fy), clr)
		return
	}
	tx, ty := math.Cos(end)*rx, math.Sin(end)*ry
	md := math.Sqrt(rx*rx + ry*ry)
	px := rx * rx / md
	py := ry * ry / md
	L, R := round(ox-px), round(ox+px)
	T, B := round(oy-py), round(oy+py)
	if (r-l)%2 == 0 {
		rx += 0.01
	}
	if (b-t)%2 == 0 {
		ry += 0.01
	}
	draw1 := func(T, B int) {
		for j := T; j <= B; j++ {
			y := float64(j) - oy
			x := math.Sqrt(1-(y*y)/(ry*ry)) * rx
			img.Set(round(ox+x), j, clr)
		}
	}
	draw2 := func(L, R int) {
		for i := L; i <= R; i++ {
			x := float64(i) - ox
			y := math.Sqrt(1-(x*x)/(rx*rx)) * ry
			img.Set(i, round(oy-y), clr)
		}
	}
	draw3 := func(T, B int) {
		for j := T; j <= B; j++ {
			y := float64(j) - oy
			x := math.Sqrt(1-(y*y)/(ry*ry)) * rx
			img.Set(round(ox-x), j, clr)
		}
	}
	draw4 := func(L, R int) {
		for i := L; i <= R; i++ {
			x := float64(i) - ox
			y := math.Sqrt(1-(x*x)/(rx*rx)) * ry
			img.Set(i, round(oy+y), clr)
		}
	}
	var p, q int
	if fx >= -px && fx <= px {
		if fy > 0 {
			p = 2
		} else {
			p = 4
		}
	} else {
		if fx > 0 {
			p = 1
		} else {
			p = 3
		}
	}
	if tx >= -px && tx <= px {
		if ty > 0 {
			q = 2
		} else {
			q = 4
		}
	} else {
		if tx > 0 {
			q = 1
		} else {
			q = 3
		}
	}
	switch p {
	case 1:
		switch q {
		case 1:
			if fy < ty {
				draw1(round(oy-ty), round(oy-fy))
			} else {
				draw1(T, round(oy-fy))
				draw2(L, R)
				draw3(T, B)
				draw4(L, R)
				draw1(round(oy-ty), B)
			}
		case 2:
			draw1(T, round(oy-fy))
			draw2(round(ox+tx), R)
		case 3:
			draw1(T, round(oy-fy))
			draw2(L, R)
			draw3(T, round(oy-ty))
		case 4:
			draw1(T, round(oy-fy))
			draw2(L, R)
			draw3(T, B)
			draw4(L, round(ox+tx))
		}
	case 2:
		switch q {
		case 1:
			draw2(L, round(ox+fx))
			draw3(T, B)
			draw4(L, R)
			draw1(round(oy-ty), B)
		case 2:
			if fx > tx {
				draw2(round(ox+tx), round(ox+fx))
			} else {
				draw2(L, round(ox+fx))
				draw3(T, B)
				draw4(L, R)
				draw1(T, B)
				draw2(round(ox+tx), R)
			}
		case 3:
			draw2(L, round(ox+fx))
			draw3(T, round(oy-ty))
		case 4:
			draw2(L, round(ox+fx))
			draw3(T, B)
			draw4(L, round(ox+tx))
		}
	case 3:
		switch q {
		case 1:
			draw3(round(oy-fy), B)
			draw4(L, R)
			draw1(round(oy-ty), B)
		case 2:
			draw3(round(oy-fy), B)
			draw4(L, R)
			draw1(T, B)
			draw2(round(ox+tx), R)
		case 3:
			if fy > ty {
				draw3(round(oy-fy), round(oy-ty))
			} else {
				draw3(round(oy-fy), B)
				draw4(L, R)
				draw1(T, B)
				draw2(L, R)
				draw3(T, round(oy-ty))
			}
		case 4:
			draw3(round(oy-fy), B)
			draw4(L, round(ox+tx))
		}
	case 4:
		switch q {
		case 1:
			draw4(round(ox+fx), R)
			draw1(round(oy-ty), B)
		case 2:
			draw4(round(ox+fx), R)
			draw1(T, B)
			draw2(round(ox+tx), R)
		case 3:
			draw4(round(ox+fx), R)
			draw1(T, B)
			draw2(L, R)
			draw3(T, round(oy-ty))
		case 4:
			if fx < tx {
				draw4(round(ox+fx), round(ox+tx))
			} else {
				draw4(round(ox+fx), R)
				draw1(T, B)
				draw2(L, R)
				draw3(T, B)
				draw4(L, round(ox+tx))
			}
		}
	}
}

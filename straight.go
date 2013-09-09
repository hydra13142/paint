package paint

import "image/color"

func Line(img Image, x1, y1, x2, y2 int, clr color.Color) {
	if x1 == x2 && y1 == y2 {
		img.Set(x1, y1, clr)
		return
	}
	abs := func(x int) int {
		if x >= 0 {
			return x
		}
		return -x
	}
	dx := x2 - x1
	dy := y2 - y1
	if abs(dx) > abs(dy) {
		i, k := 0, 0
		if dx > 0 {
			i = 1
		} else {
			i = -1
		}
		for ; x1 != x2; x1 += i {
			img.Set(x1, k*dy/dx+y1, clr)
			k += i
		}
	} else {
		j, k := 0, 0
		if dy > 0 {
			j = 1
		} else {
			j = -1
		}
		for ; y1 != y2; y1 += j {
			img.Set(k*dx/dy+x1, y1, clr)
			k += j
		}
	}
}

func Rect(img Image, l, t, r, b int, clr color.Color) {
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
	for x := l; x <= r; x++ {
		img.Set(x, t, clr)
		img.Set(x, b, clr)
	}
	for y := t; y <= b; y++ {
		img.Set(l, y, clr)
		img.Set(r, y, clr)
	}
}

func Block(img Image, l, t, r, b int, clr color.Color) {
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
	for x := l; x <= r; x++ {
		for y := t; y <= b; y++ {
			img.Set(x, y, clr)
		}
	}
}

func Region(img Image, l, t, r, b int, f NetType, d int, clr color.Color) {
	var x, y, k int
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
	for x = l; x <= r; x++ {
		img.Set(x, t, clr)
		img.Set(x, b, clr)
	}
	for y = t; y <= b; y++ {
		img.Set(l, y, clr)
		img.Set(r, y, clr)
	}
	if f&Level != 0 {
		for y = t + d; y < b; y += d {
			Line(img, l, y, r, y, clr)
		}
	}
	if f&Plumb != 0 {
		for x = l + d; x < r; x += d {
			Line(img, x, t, x, b, clr)
		}
	}
	if f&Slant != 0 {
		for y = t; y < b; y += d {
			k = r - l
			if b-y < k {
				k = b - y
			}
			Line(img, l, y, l+k, y+k, clr)
		}
		for x = l + d; x < r; x += d {
			k = b - t
			if r-x < k {
				k = r - x
			}
			Line(img, x, t, x+k, t+k, clr)
		}
	}
	if f&Twill != 0 {
		for y = b; y > t; y -= d {
			k = r - l
			if y-t < k {
				k = y - t
			}
			Line(img, l, y, l+k, y-k, clr)
		}
		for x = l + d; x < r; x += d {
			k = b - t
			if r-x < k {
				k = r - x
			}
			Line(img, x, b, x+k, b-k, clr)
		}
	}
}

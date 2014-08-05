package paint

// 绘制点
func (img *Image) Pset(x, y int) {
	img.Set(x, y, img.FR)
}

// 绘制直线
func (img *Image) Line(x1, y1, x2, y2 int) {
	if x1 == x2 && y1 == y2 {
		img.Pset(x1, y1)
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
			img.Pset(x1, k*dy/dx+y1)
			k += i
		}
		img.Pset(x2, y2)
	} else {
		j, k := 0, 0
		if dy > 0 {
			j = 1
		} else {
			j = -1
		}
		for ; y1 != y2; y1 += j {
			img.Pset(k*dx/dy+x1, y1)
			k += j
		}
		img.Pset(x2, y2)
	}
}

// 绘制矩形
func (img *Image) Rect(l, t, r, b int) {
	if l == r || t == b {
		img.Line(l, t, r, b)
		return
	}
	if l > r {
		l, r = r, l
	}
	if t > b {
		t, b = b, t
	}
	for x := l; x <= r; x++ {
		img.Pset(x, t)
		img.Pset(x, b)
	}
	for y := t; y <= b; y++ {
		img.Pset(l, y)
		img.Pset(r, y)
	}
}

// 绘制矩形并填充
func (img *Image) Bar(l, t, r, b int) {
	if l == r || t == b {
		img.Line(l, t, r, b)
		return
	}
	if l > r {
		l, r = r, l
	}
	if t > b {
		t, b = b, t
	}
	img.Line(l, t, r, t)
	img.Line(l, b, r, b)
	img.Line(l, t, l, b)
	img.Line(r, t, r, b)
	for x := l + 1; x < r; x++ {
		for y := t + 1; y < b; y++ {
			img.Set(x, y, img.BG)
		}
	}
}

// 网格的类型
type NetType int

const (
	Level NetType = 1 << iota // 水平
	Plumb                     // 竖直
	Slant                     // 斜杠
	Twill                     // 反斜杠
)

// 绘制矩形并在内部绘制网格
func (img *Image) Block(l, t, r, b int, f NetType, d int) {
	var x, y, k int
	if l == r || t == b {
		img.Line(l, t, r, b)
		return
	}
	if l > r {
		l, r = r, l
	}
	if t > b {
		t, b = b, t
	}
	for x = l; x <= r; x++ {
		img.Pset(x, t)
		img.Pset(x, b)
	}
	for y = t; y <= b; y++ {
		img.Pset(l, y)
		img.Pset(r, y)
	}
	if f&Level != 0 {
		for y = t + d; y < b; y += d {
			img.Line(l, y, r, y)
		}
	}
	if f&Plumb != 0 {
		for x = l + d; x < r; x += d {
			img.Line(x, t, x, b)
		}
	}
	if f&Slant != 0 {
		for y = t; y < b; y += d {
			k = r - l
			if b-y < k {
				k = b - y
			}
			img.Line(l, y, l+k, y+k)
		}
		for x = l + d; x < r; x += d {
			k = b - t
			if r-x < k {
				k = r - x
			}
			img.Line(x, t, x+k, t+k)
		}
	}
	if f&Twill != 0 {
		for y = b; y > t; y -= d {
			k = r - l
			if y-t < k {
				k = y - t
			}
			img.Line(l, y, l+k, y-k)
		}
		for x = l + d; x < r; x += d {
			k = b - t
			if r-x < k {
				k = r - x
			}
			img.Line(x, b, x+k, b-k)
		}
	}
}

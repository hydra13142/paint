package paint

func (c *Canvas) Line(x1, y1, x2, y2 int) {
	if x1 == x2 && y1 == y2 {
		c.Image.Set(x1, y1, c.Fore)
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
			c.Image.Set(x1, k*dy/dx+y1, c.Fore)
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
			c.Image.Set(k*dx/dy+x1, y1, c.Fore)
			k += j
		}
	}
}

func (c *Canvas) Rect(l, t, r, b int) {
	if l == r || t == b {
		c.Line(l, t, r, b)
		return
	}
	if l > r {
		l, r = r, l
	}
	if t > b {
		t, b = b, t
	}
	for x := l; x <= r; x++ {
		c.Image.Set(x, t, c.Fore)
		c.Image.Set(x, b, c.Fore)
	}
	for y := t; y <= b; y++ {
		c.Image.Set(l, y, c.Fore)
		c.Image.Set(r, y, c.Fore)
	}
}

func (c *Canvas) Block(l, t, r, b int) {
	if l == r || t == b {
		c.Line(l, t, r, b)
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
			c.Image.Set(x, y, c.Fore)
		}
	}
}

func (c *Canvas) Region(l, t, r, b int, f NetType, d int) {
	var x, y, k int
	if l == r || t == b {
		c.Line(l, t, r, b)
		return
	}
	if l > r {
		l, r = r, l
	}
	if t > b {
		t, b = b, t
	}
	for x = l; x <= r; x++ {
		c.Image.Set(x, t, c.Fore)
		c.Image.Set(x, b, c.Fore)
	}
	for y = t; y <= b; y++ {
		c.Image.Set(l, y, c.Fore)
		c.Image.Set(r, y, c.Fore)
	}
	if f&Level != 0 {
		for y = t + d; y < b; y += d {
			c.Line(l, y, r, y)
		}
	}
	if f&Plumb != 0 {
		for x = l + d; x < r; x += d {
			c.Line(x, t, x, b)
		}
	}
	if f&Slant != 0 {
		for y = t; y < b; y += d {
			k = r - l
			if b-y < k {
				k = b - y
			}
			c.Line(l, y, l+k, y+k)
		}
		for x = l + d; x < r; x += d {
			k = b - t
			if r-x < k {
				k = r - x
			}
			c.Line(x, t, x+k, t+k)
		}
	}
	if f&Twill != 0 {
		for y = b; y > t; y -= d {
			k = r - l
			if y-t < k {
				k = y - t
			}
			c.Line(l, y, l+k, y-k)
		}
		for x = l + d; x < r; x += d {
			k = b - t
			if r-x < k {
				k = r - x
			}
			c.Line(x, b, x+k, b-k)
		}
	}
}

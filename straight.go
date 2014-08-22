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
func (img *Image) Rect(smlX, smlY, bigX, bigY int) {
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
	for x := smlX; x <= bigX; x++ {
		img.Pset(x, smlY)
		img.Pset(x, bigY)
	}
	for y := smlY; y <= bigY; y++ {
		img.Pset(smlX, y)
		img.Pset(bigX, y)
	}
}

// 绘制矩形并填充
func (img *Image) Bar(smlX, smlY, bigX, bigY int) {
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
	img.Line(smlX, smlY, bigX, smlY)
	img.Line(smlX, smlY, smlX, bigY)
	img.Line(smlX, bigY, bigX, bigY)
	img.Line(bigX, smlY, bigX, bigY)
	for x := smlX + 1; x < bigX; x++ {
		for y := smlY + 1; y < bigY; y++ {
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
func (img *Image) Block(smlX, smlY, bigX, bigY int, f NetType, d int) {
	var x, y, k int
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
	for x = smlX; x <= bigX; x++ {
		img.Pset(x, smlY)
		img.Pset(x, bigY)
	}
	for y = smlY; y <= bigY; y++ {
		img.Pset(smlX, y)
		img.Pset(bigX, y)
	}
	if f&Level != 0 {
		for y = smlY + d; y < bigY; y += d {
			img.Line(smlX, y, bigX, y)
		}
	}
	if f&Plumb != 0 {
		for x = smlX + d; x < bigX; x += d {
			img.Line(x, smlY, x, bigY)
		}
	}
	if f&Slant != 0 {
		for y = smlY; y < bigY; y += d {
			k = bigX - smlX
			if bigY-y < k {
				k = bigY - y
			}
			img.Line(smlX, y, smlX+k, y+k)
		}
		for x = smlX + d; x < bigX; x += d {
			k = bigY - smlY
			if bigX-x < k {
				k = bigX - x
			}
			img.Line(x, smlY, x+k, smlY+k)
		}
	}
	if f&Twill != 0 {
		for y = bigY; y > smlY; y -= d {
			k = bigX - smlX
			if y-smlY < k {
				k = y - smlY
			}
			img.Line(smlX, y, smlX+k, y-k)
		}
		for x = smlX + d; x < bigX; x += d {
			k = bigY - smlY
			if bigX-x < k {
				k = bigX - x
			}
			img.Line(x, bigY, x+k, bigY-k)
		}
	}
}

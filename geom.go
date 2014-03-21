package paint

import (
	"math"
	"sort"
)

const (
	Inner int = 1 << iota
	Edge
	Extend
	Left
	Right
	Inside  = Inner | Edge
	Outside = Extend | Left | Right
)

type Point struct {
	X, Y float64
}

type Vector struct {
	I, J float64
}

type Segment struct {
	A, B Point
}

type Convex []Point // 顺时针给出的凸多边形的各个顶点

type points struct {
	avr Point
	spt []Point
}

func (s Segment) Place(p Point) int {
	pax, pay := p.X-s.A.X, p.Y-s.A.Y
	pbx, pby := p.X-s.B.X, p.Y-s.B.Y
	switch k := pbx*pay - pax*pby; {
	case k < 0:
		return Left
	case k > 0:
		return Right
	default:
		if pax*pbx+pay*pby <= 0 {
			return Inside
		}
	}
	return Extend
}

func (s Segment) Cross(l Segment) ([]Point, int) {
	if s.A == s.B {
		if l.Place(s.A) == Inside {
			return []Point{s.A}, 1
		}
		return nil, 0
	}
	if l.A == l.B {
		if s.Place(l.A) == Inside {
			return []Point{l.A}, 1
		}
		return nil, 0
	}
	V1 := Vector{s.B.X - s.A.X, s.B.Y - s.A.Y}
	V2 := Vector{l.B.X - l.A.X, l.B.Y - l.A.Y}
	V3 := Vector{l.A.X - s.A.X, l.A.Y - s.A.Y}
	d := V1.I*V2.J - V2.I*V1.J
	if d != 0 {
		k1 := (V3.I*V2.J - V2.I*V3.J) / d
		if k1 < 0 || k1 > 1 {
			return nil, 0
		}
		k2 := (V3.I*V1.J - V1.I*V3.J) / d
		if k2 < 0 || k2 > 1 {
			return nil, 0
		}
		return []Point{{s.A.X + V1.I*k1, s.A.Y + V1.J*k1}}, 1
	}
	if V3.I*V2.J-V2.I*V3.J != 0 {
		return nil, 0
	}
	var la, lb float64
	if math.Abs(V1.I) > math.Abs(V1.J) {
		la = (l.A.X - s.A.X) / V1.I
		lb = (l.B.X - s.A.X) / V1.I
	} else {
		la = (l.A.Y - s.A.Y) / V1.J
		lb = (l.B.Y - s.A.Y) / V1.J
	}
	if la < lb {
		switch {
		case la > 1:
			return nil, 0
		case la == 1:
			return []Point{l.A}, 1
		case la >= 0:
			if lb > 1 {
				return []Point{l.A, s.B}, 2
			} else {
				return []Point{l.A, l.B}, 2
			}
		default:
			switch {
			case lb > 1:
				return []Point{s.A, s.B}, 2
			case lb > 0:
				return []Point{s.A, l.B}, 2
			case lb == 0:
				return []Point{l.B}, 1
			default:
				return nil, 0
			}
		}
	} else {
		switch {
		case lb > 1:
			return nil, 0
		case lb == 1:
			return []Point{l.B}, 1
		case lb >= 0:
			if la > 1 {
				return []Point{s.B, l.B}, 2
			} else {
				return []Point{l.A, l.B}, 2
			}
		default:
			switch {
			case la > 1:
				return []Point{s.B, s.A}, 2
			case la > 0:
				return []Point{l.A, s.A}, 2
			case la == 0:
				return []Point{l.A}, 1
			default:
				return nil, 0
			}
		}
	}
}

func (s Segment) Length() float64 {
	dx := s.B.X - s.A.X
	dy := s.B.Y - s.A.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func (c Convex) Place(p Point) int {
	for i, l := 0, len(c); i < l; i++ {
		s := Segment{c[i], c[(i+1)%l]}
		switch s.Place(p) {
		case Left:
			return Outside
		case Inside:
			return Edge
		}
	}
	return Inner
}

func (c Convex) Area() float64 {
	area := func(A, B, C Point) float64 {
		V1 := Vector{B.X - A.X, B.Y - A.Y}
		V2 := Vector{C.X - A.X, C.Y - A.Y}
		return (V2.I*V1.J - V1.I*V2.J) / 2
	}
	s := 0.0
	for i := 2; i < len(c); i++ {
		s += area(c[0], c[i-1], c[i])
	}
	return s
}

func (s *points) Len() int {
	return len(s.spt)
}

func (s *points) Less(i, j int) bool {
	ix, iy := s.spt[i].X-s.avr.X, s.spt[i].Y-s.avr.Y
	jx, jy := s.spt[j].X-s.avr.X, s.spt[j].Y-s.avr.Y
	return math.Atan2(iy, ix) > math.Atan2(jy, jx)
}

func (s *points) Swap(i, j int) {
	s.spt[i], s.spt[j] = s.spt[j], s.spt[i]
}

func (s *points) Average() {
	l := len(s.spt)
	if l == 0 {
		s.avr.X, s.avr.Y = 0, 0
	}
	x, y := 0.0, 0.0
	for _, p := range s.spt {
		x += p.X
		y += p.Y
	}
	s.avr.X, s.avr.Y = x/float64(l), y/float64(l)
}

func (s *points) Light() {
	if len(s.spt) <= 1 {
		return
	}
	n := make([]Point, 1, len(s.spt))
	e := s.spt[0]
	n[0] = e
	for _, p := range s.spt[1:] {
		if p != e {
			n = append(n, p)
			e = p
		}
	}
	s.spt = n
}

func And(a, b Convex) Convex {
	c := []Point{}
	for _, p := range a {
		if b.Place(p) == Inner {
			c = append(c, p)
		}
	}
	for _, p := range b {
		if a.Place(p) == Inner {
			c = append(c, p)
		}
	}
	for i, la := 0, len(a); i < la; i++ {
		ni := (i + 1) % la
		for j, lb := 0, len(b); j < lb; j++ {
			nj := (j + 1) % lb
			x := Segment{a[i], a[ni]}
			y := Segment{b[j], b[nj]}
			if r, t := x.Cross(y); t == 1 {
				c = append(c, r[0])
			}
		}
	}
	if len(c) < 3 {
		return nil
	}
	ts := &points{spt: c}
	ts.Average()
	sort.Sort(ts)
	ts.Light()
	if c = ts.spt; len(c) < 3 {
		return nil
	}
	return c
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func NewSegment(x1, y1, x2, y2 float64) Segment {
	return Segment{Point{x1, y1}, Point{x2, y2}}
}

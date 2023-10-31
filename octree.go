package octree

type PointCloud interface {
	Insert(point *PointData)
	Select(minpoint, maxpoint Point) []*PointData
}

type octree struct {
	root *Box
}

type Point struct {
	X, Y, Z float64
}

type PointData struct {
	Point
}

type Box struct {
	minpoint Point
	maxpoint Point
	midpoint Point
	deep     int
	count    int
	maxcount int
	children []*Box
	points   []*PointData
}

func NewOctree(minpoint, maxpoint Point, maxdeep int, maxcount int) PointCloud {
	return &octree{root: NewBox(minpoint, maxpoint, maxdeep, maxcount)}
}

func NewBox(minpoint, maxpoint Point, deep int, maxcount int) *Box {
	return &Box{minpoint, maxpoint, Point{
		X: (minpoint.X + maxpoint.X) / 2,
		Y: (minpoint.Y + maxpoint.Y) / 2,
		Z: (minpoint.Z + maxpoint.Z) / 2,
	}, deep, 0, maxcount, nil, make([]*PointData, 0)}
}

func (o *octree) Insert(point *PointData) {
	o.root.Insert(point)
}

func (o *octree) Select(minpoint, maxpoint Point) []*PointData {
	return o.root.Select(minpoint, maxpoint)
}

func (b *Box) Select(minpoint, maxpoint Point) []*PointData {
	result := make([]*PointData, 0)
	if b.maxpoint.X < minpoint.X || b.maxpoint.Y < minpoint.Y || b.maxpoint.Z < minpoint.Z ||
		b.minpoint.X > maxpoint.X || b.minpoint.Y > maxpoint.Y || b.minpoint.Z > maxpoint.Z {
		return result
	}
	if b.children == nil {
		for _, point := range b.points {
			if point.X >= minpoint.X && point.Y >= minpoint.Y && point.Z >= minpoint.Z &&
				point.X <= maxpoint.X && point.Y <= maxpoint.Y && point.Z <= maxpoint.Z {
				result = append(result, point)
			}
		}
	} else {
		for _, c := range b.children {
			result = append(result, c.Select(minpoint, maxpoint)...)
		}
	}
	return result
}

func (b *Box) Insert(point *PointData) {
	if b.children == nil {
		b.points = append(b.points, point)
		if b.deep > 0 && len(b.points) > b.maxcount {
			b.expand()
			for _, p := range b.points {
				b.children[b.ocpos(p)].Insert(p)
			}
			b.points = make([]*PointData, 0)
		}
	} else {
		b.children[b.ocpos(point)].Insert(point)
	}
	b.count++
}

/*
*
	for i := range b.children {
		mn := Point{b.minpoint.X, b.minpoint.Y, b.minpoint.Z}
		mx := Point{b.maxpoint.X, b.maxpoint.Y, b.maxpoint.Z}
		if i&1 == 1 {
			mn.X = b.midpoint.X
		} else {
			mx.X = b.midpoint.X
		}
		if i&2 == 2 {
			mn.Y = b.midpoint.Y
		} else {
			mx.Y = b.midpoint.Y
		}
		if i&4 == 4 {
			mn.Z = b.midpoint.Z
		} else {
			mx.Z = b.midpoint.Z
		}
		b.children[i] = NewBox(mn, mx, b.deep-1, b.maxcount)
	}
*/
func (b *Box) expand() {
	b.children = []*Box{
		NewBox(Point{b.minpoint.X, b.minpoint.Y, b.minpoint.Z}, Point{b.midpoint.X, b.midpoint.Y, b.midpoint.Z}, b.deep-1, b.maxcount),
		NewBox(Point{b.midpoint.X, b.minpoint.Y, b.minpoint.Z}, Point{b.maxpoint.X, b.midpoint.Y, b.midpoint.Z}, b.deep-1, b.maxcount),
		NewBox(Point{b.minpoint.X, b.midpoint.Y, b.minpoint.Z}, Point{b.midpoint.X, b.maxpoint.Y, b.midpoint.Z}, b.deep-1, b.maxcount),
		NewBox(Point{b.midpoint.X, b.midpoint.Y, b.minpoint.Z}, Point{b.maxpoint.X, b.maxpoint.Y, b.midpoint.Z}, b.deep-1, b.maxcount),
		NewBox(Point{b.minpoint.X, b.minpoint.Y, b.midpoint.Z}, Point{b.midpoint.X, b.midpoint.Y, b.maxpoint.Z}, b.deep-1, b.maxcount),
		NewBox(Point{b.midpoint.X, b.minpoint.Y, b.midpoint.Z}, Point{b.maxpoint.X, b.midpoint.Y, b.maxpoint.Z}, b.deep-1, b.maxcount),
		NewBox(Point{b.minpoint.X, b.midpoint.Y, b.midpoint.Z}, Point{b.midpoint.X, b.maxpoint.Y, b.maxpoint.Z}, b.deep-1, b.maxcount),
		NewBox(Point{b.midpoint.X, b.midpoint.Y, b.midpoint.Z}, Point{b.maxpoint.X, b.maxpoint.Y, b.maxpoint.Z}, b.deep-1, b.maxcount),
	}
}

func (b *Box) ocpos(point *PointData) int {
	p := 0
	if b.midpoint.X < point.X {
		p += 1
	}
	if b.midpoint.Y < point.Y {
		p += 2
	}
	if b.midpoint.Z < point.Z {
		p += 4
	}
	return p
}

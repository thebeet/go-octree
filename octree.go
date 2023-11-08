package octree

type PointCloud interface {
	Insert(point []*PointData)
	Select(minpoint, maxpoint Point) []*PointData
}

type octree struct {
	root          *Box
	points        []*PointData
	maxPointCount int
	maxDeep       int
}

type Point struct {
	X, Y, Z float64
}

type PointData struct {
	Point
	// payload data
}

type Box struct {
	minpoint Point
	maxpoint Point
	midpoint Point
	children []*Box
	points   []*PointData
}

func NewOctree(minpoint, maxpoint Point, maxDeep int, maxPointCount int) PointCloud {
	return &octree{
		root:          NewBox(minpoint, maxpoint),
		points:        nil,
		maxPointCount: maxPointCount,
		maxDeep:       maxDeep,
	}
}

func NewBox(minpoint, maxpoint Point) *Box {
	return &Box{minpoint, maxpoint, Point{
		X: (minpoint.X + maxpoint.X) / 2,
		Y: (minpoint.Y + maxpoint.Y) / 2,
		Z: (minpoint.Z + maxpoint.Z) / 2,
	}, nil, nil}
}

func (o *octree) Insert(points []*PointData) {
	o.points = make([]*PointData, len(points))
	copy(o.points, points)
	o.root.points = o.points
	o.root.split(o.maxPointCount, o.maxDeep)
}

func cmpx(p1, p2 *Point) bool {
	return p1.X < p2.X
}

func cmpy(p1, p2 *Point) bool {
	return p1.Y < p2.Y
}

func cmpz(p1, p2 *Point) bool {
	return p1.Z < p2.Z
}

func (b *Box) split(maxPointCount int, deep int) {
	if (len(b.points) > maxPointCount) && (deep > 0) {
		p := make([]int, 9)
		p[0], p[8] = 0, len(b.points)
		p[4] = b.sortinner(b.points, cmpz)
		p[2], p[6] = b.sortinner(b.points[p[0]:p[4]], cmpy), b.sortinner(b.points[p[4]:p[8]], cmpy)+p[4]
		p[1], p[3], p[5], p[7] = b.sortinner(b.points[p[0]:p[2]], cmpx), b.sortinner(b.points[p[2]:p[4]], cmpx)+p[2],
			b.sortinner(b.points[p[4]:p[6]], cmpx)+p[4], b.sortinner(b.points[p[6]:p[8]], cmpx)+p[6]
		b.children = []*Box{
			NewBox(Point{b.minpoint.X, b.minpoint.Y, b.minpoint.Z}, Point{b.midpoint.X, b.midpoint.Y, b.midpoint.Z}),
			NewBox(Point{b.midpoint.X, b.minpoint.Y, b.minpoint.Z}, Point{b.maxpoint.X, b.midpoint.Y, b.midpoint.Z}),
			NewBox(Point{b.minpoint.X, b.midpoint.Y, b.minpoint.Z}, Point{b.midpoint.X, b.maxpoint.Y, b.midpoint.Z}),
			NewBox(Point{b.midpoint.X, b.midpoint.Y, b.minpoint.Z}, Point{b.maxpoint.X, b.maxpoint.Y, b.midpoint.Z}),
			NewBox(Point{b.minpoint.X, b.minpoint.Y, b.midpoint.Z}, Point{b.midpoint.X, b.midpoint.Y, b.maxpoint.Z}),
			NewBox(Point{b.midpoint.X, b.minpoint.Y, b.midpoint.Z}, Point{b.maxpoint.X, b.midpoint.Y, b.maxpoint.Z}),
			NewBox(Point{b.minpoint.X, b.midpoint.Y, b.midpoint.Z}, Point{b.midpoint.X, b.maxpoint.Y, b.maxpoint.Z}),
			NewBox(Point{b.midpoint.X, b.midpoint.Y, b.midpoint.Z}, Point{b.maxpoint.X, b.maxpoint.Y, b.maxpoint.Z}),
		}
		for i := 0; i < len(b.children); i++ {
			b.children[i].points = b.points[p[i]:p[i+1]]
			b.children[i].split(maxPointCount, deep-1)
		}
	}
}

func (b *Box) sortinner(points []*PointData, cmp func(p1, p2 *Point) bool) int {
	f, t := 0, len(points)-1
	for {
		for (f <= t) && cmp(&points[f].Point, &b.midpoint) {
			f++
		}
		for (f <= t) && !cmp(&points[t].Point, &b.midpoint) {
			t--
		}
		if f < t {
			points[f], points[t] = points[t], points[f]
			f++
			t--
		} else {
			break
		}
	}
	return t + 1
}

func (o *octree) Select(minpoint, maxpoint Point) []*PointData {
	return o.root.Select(minpoint, maxpoint)
}

func (b *Box) Select(minpoint, maxpoint Point) []*PointData {
	if b.minpoint.X > minpoint.X && b.minpoint.Y > minpoint.Y && b.minpoint.Z > minpoint.Z &&
		b.maxpoint.X < maxpoint.X && b.maxpoint.Y < maxpoint.Y && b.maxpoint.Z < maxpoint.Z {
		return b.points
	}
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

package octree

type plain struct {
	points []*PointData
}

func NewPlain() PointCloud {
	return &plain{make([]*PointData, 0)}
}

func (p *plain) Insert(point *PointData) {
	p.points = append(p.points, point)
}

func (p *plain) Select(minpoint, maxpoint Point) []*PointData {
	result := make([]*PointData, 0)
	for _, point := range p.points {
		if point.X >= minpoint.X && point.Y >= minpoint.Y && point.Z >= minpoint.Z &&
			point.X <= maxpoint.X && point.Y <= maxpoint.Y && point.Z <= maxpoint.Z {
			result = append(result, point)
		}
	}
	return result
}

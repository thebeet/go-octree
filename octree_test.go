package octree_test

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thebeet/go-octree"
)

func TestSelect(t *testing.T) {
	assert := assert.New(t)
	points := pointData(1000000)
	cloudOctree := octree.NewOctree(octree.Point{0, 0, 0}, octree.Point{1, 1, 1}, 8, 256)
	cloudOctree.Insert(points)

	cloudPlain := octree.NewPlain()
	cloudPlain.Insert(points)

	rand := rand.New(rand.NewSource(2))
	for n := 0; n < 1000; n++ {
		mn := octree.Point{
			X: rand.Float64() * 0.9,
			Y: rand.Float64() * 0.9,
			Z: rand.Float64() * 0.9,
		}
		mx := octree.Point{
			X: mn.X + rand.Float64()*0.1,
			Y: mn.Y + rand.Float64()*0.1,
			Z: mn.Z + rand.Float64()*0.1,
		}
		ocResult := cloudOctree.Select(mn, mx)
		plainResult := cloudPlain.Select(mn, mx)
		assert.True(len(ocResult) == len(plainResult))
	}
}

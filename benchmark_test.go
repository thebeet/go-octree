package octree_test

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thebeet/go-octree"
)

func pointData(c int) []*octree.PointData {
	rand := rand.New(rand.NewSource(1))
	points := make([]*octree.PointData, c)
	for i := 0; i < c; i++ {
		points[i] = &octree.PointData{
			octree.Point{
				X: rand.Float64(),
				Y: rand.Float64(),
				Z: rand.Float64(),
			},
		}
	}
	return points
}

func benchmarkOctreeInit(b *testing.B, deep, maxcount int) {
	assert := assert.New(b)

	// init test data
	points := pointData(5000000)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		tree := octree.NewOctree(octree.Point{0, 0, 0}, octree.Point{1, 1, 1}, deep, maxcount)
		assert.NotNil(tree)
		for _, p := range points {
			tree.Insert(p)
		}
	}
}
func BenchmarkOctreeInit5000000_8_1024(b *testing.B) {
	benchmarkOctreeInit(b, 8, 1024)
}
func BenchmarkOctreeInit5000000_8_512(b *testing.B) {
	benchmarkOctreeInit(b, 8, 512)
}
func BenchmarkOctreeInit5000000_8_256(b *testing.B) {
	benchmarkOctreeInit(b, 8, 256)
}
func BenchmarkOctreeInit5000000_8_128(b *testing.B) {
	benchmarkOctreeInit(b, 8, 128)
}
func BenchmarkOctreeInit5000000_8_64(b *testing.B) {
	benchmarkOctreeInit(b, 8, 64)
}
func BenchmarkOctreeInit5000000_8_32(b *testing.B) {
	benchmarkOctreeInit(b, 8, 32)
}

func BenchmarkOctreeInit5000000_9_1024(b *testing.B) {
	benchmarkOctreeInit(b, 9, 1024)
}
func BenchmarkOctreeInit5000000_9_512(b *testing.B) {
	benchmarkOctreeInit(b, 9, 512)
}
func BenchmarkOctreeInit5000000_9_256(b *testing.B) {
	benchmarkOctreeInit(b, 9, 256)
}
func BenchmarkOctreeInit5000000_9_128(b *testing.B) {
	benchmarkOctreeInit(b, 9, 128)
}
func BenchmarkOctreeInit5000000_9_64(b *testing.B) {
	benchmarkOctreeInit(b, 9, 64)
}
func BenchmarkOctreeInit5000000_9_32(b *testing.B) {
	benchmarkOctreeInit(b, 9, 32)
}

func BenchmarkOctreeInit5000000_10_1024(b *testing.B) {
	benchmarkOctreeInit(b, 10, 1024)
}
func BenchmarkOctreeInit5000000_10_512(b *testing.B) {
	benchmarkOctreeInit(b, 10, 512)
}
func BenchmarkOctreeInit5000000_10_256(b *testing.B) {
	benchmarkOctreeInit(b, 10, 256)
}
func BenchmarkOctreeInit5000000_10_128(b *testing.B) {
	benchmarkOctreeInit(b, 10, 128)
}
func BenchmarkOctreeInit5000000_10_64(b *testing.B) {
	benchmarkOctreeInit(b, 10, 64)
}
func BenchmarkOctreeInit5000000_10_32(b *testing.B) {
	benchmarkOctreeInit(b, 9, 32)
}

func BenchmarkPlainInit5000000(b *testing.B) {
	assert := assert.New(b)

	// init test data
	points := pointData(5000000)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		tree := octree.NewPlain()
		assert.NotNil(tree)
		for _, p := range points {
			tree.Insert(p)
		}
	}
}

func benchOctreeSelect(b *testing.B, deep, maxcount int) {
	// init test data
	points := pointData(5000000)
	cloud := octree.NewOctree(octree.Point{0, 0, 0}, octree.Point{1, 1, 1}, deep, maxcount)
	for _, p := range points {
		cloud.Insert(p)
	}
	b.ResetTimer()

	rand := rand.New(rand.NewSource(2))
	for n := 0; n < b.N; n++ {
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
		cloud.Select(mn, mx)
	}
}
func BenchmarkOctreeSelect_8_65536_5000000(b *testing.B) {
	benchOctreeSelect(b, 8, 65536)
}
func BenchmarkOctreeSelect_8_32768_5000000(b *testing.B) {
	benchOctreeSelect(b, 8, 32768)
}
func BenchmarkOctreeSelect_8_16384_5000000(b *testing.B) {
	benchOctreeSelect(b, 8, 16384)
}
func BenchmarkOctreeSelect_8_8192_5000000(b *testing.B) {
	benchOctreeSelect(b, 8, 8192)
}
func BenchmarkOctreeSelect_8_4096_5000000(b *testing.B) {
	benchOctreeSelect(b, 8, 4096)
}
func BenchmarkOctreeSelect_8_2048_5000000(b *testing.B) {
	benchOctreeSelect(b, 8, 2048)
}
func BenchmarkOctreeSelect_8_1024_5000000(b *testing.B) {
	benchOctreeSelect(b, 8, 1024)
}
func BenchmarkOctreeSelect_8_512_5000000(b *testing.B) {
	benchOctreeSelect(b, 8, 512)
}
func BenchmarkOctreeSelect_8_256_5000000(b *testing.B) {
	benchOctreeSelect(b, 8, 256)
}
func BenchmarkOctreeSelect_8_128_5000000(b *testing.B) {
	benchOctreeSelect(b, 8, 128)
}
func BenchmarkOctreeSelect_8_64_5000000(b *testing.B) {
	benchOctreeSelect(b, 8, 64)
}
func BenchmarkOctreeSelect_8_32_5000000(b *testing.B) {
	benchOctreeSelect(b, 8, 32)
}
func BenchmarkOctreeSelect_8_16_5000000(b *testing.B) {
	benchOctreeSelect(b, 8, 16)
}
func BenchmarkOctreeSelect_8_8_5000000(b *testing.B) {
	benchOctreeSelect(b, 8, 8)
}
func BenchmarkOctreeSelect_9_4096_5000000(b *testing.B) {
	benchOctreeSelect(b, 9, 4096)
}
func BenchmarkOctreeSelect_9_2048_5000000(b *testing.B) {
	benchOctreeSelect(b, 9, 2048)
}
func BenchmarkOctreeSelect_9_1024_5000000(b *testing.B) {
	benchOctreeSelect(b, 9, 1024)
}
func BenchmarkOctreeSelect_9_512_5000000(b *testing.B) {
	benchOctreeSelect(b, 9, 512)
}
func BenchmarkOctreeSelect_9_256_5000000(b *testing.B) {
	benchOctreeSelect(b, 9, 256)
}
func BenchmarkOctreeSelect_9_128_5000000(b *testing.B) {
	benchOctreeSelect(b, 9, 128)
}
func BenchmarkOctreeSelect_9_64_5000000(b *testing.B) {
	benchOctreeSelect(b, 9, 64)
}
func BenchmarkOctreeSelect_9_32_5000000(b *testing.B) {
	benchOctreeSelect(b, 9, 32)
}
func BenchmarkOctreeSelect_9_16_5000000(b *testing.B) {
	benchOctreeSelect(b, 9, 16)
}
func BenchmarkOctreeSelect_10_4096_5000000(b *testing.B) {
	benchOctreeSelect(b, 10, 4096)
}
func BenchmarkOctreeSelect_10_2048_5000000(b *testing.B) {
	benchOctreeSelect(b, 10, 2048)
}
func BenchmarkOctreeSelect_10_1024_5000000(b *testing.B) {
	benchOctreeSelect(b, 10, 1024)
}
func BenchmarkOctreeSelect_10_512_5000000(b *testing.B) {
	benchOctreeSelect(b, 10, 512)
}
func BenchmarkOctreeSelect_10_256_5000000(b *testing.B) {
	benchOctreeSelect(b, 10, 256)
}
func BenchmarkOctreeSelect_10_128_5000000(b *testing.B) {
	benchOctreeSelect(b, 10, 128)
}
func BenchmarkOctreeSelect_10_64_5000000(b *testing.B) {
	benchOctreeSelect(b, 10, 64)
}
func BenchmarkOctreeSelect_10_32_5000000(b *testing.B) {
	benchOctreeSelect(b, 10, 32)
}
func BenchmarkOctreeSelect_10_16_5000000(b *testing.B) {
	benchOctreeSelect(b, 10, 16)
}
func BenchmarkPlainSelect5000000(b *testing.B) {
	// init test data
	points := pointData(5000000)
	cloud := octree.NewPlain()
	for _, p := range points {
		cloud.Insert(p)
	}
	b.ResetTimer()

	rand := rand.New(rand.NewSource(2))
	for n := 0; n < b.N; n++ {
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
		cloud.Select(mn, mx)
	}
}

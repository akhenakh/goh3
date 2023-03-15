package h3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertLatLng(t *testing.T, expected, actual LatLng) {
	almostEqual(t, expected.Lat, actual.Lat, "latitude mismatch")
	almostEqual(t, expected.Lng, actual.Lng, "longitude mismatch")
}

func TestBatchFromGeo(t *testing.T) {
	t.Parallel()

	b := NewBatch()
	defer b.Close()

	h := b.LatLngToCell(LatLng{
		Lat: 67.1509268640,
		Lng: -168.3908885810,
	}, 5)
	assert.Equal(t, validCell, h, "expected %x but got %x", validCell, h)
	assert.Equal(t, validCell, h)
}

func TestBatchToGeo(t *testing.T) {
	t.Parallel()

	b := NewBatch()
	defer b.Close()
	g := b.CellToLatLng(validCell)
	assertLatLng(t, validLatLng1, g)
}

func BenchmarkBatchToGeo(b *testing.B) {
	c := NewBatch()
	defer c.Close()

	for i := 0; i < b.N; i++ {
	}
}

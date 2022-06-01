package h3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertFromGeo(t *testing.T) {
	t.Parallel()

	c := NewConvert()
	defer c.Close()

	h := c.FromGeo(GeoCoord{
		Latitude:  67.194013596,
		Longitude: 191.598258018,
	}, 5)
	assert.Equal(t, validH3Index, h, "expected %x but got %x", validH3Index, h)
	assert.Equal(t, validH3Index, h)
}

func TestConvertToGeo(t *testing.T) {
	t.Parallel()

	c := NewConvert()
	defer c.Close()

	g := c.ToGeo(validH3Index)
	assertGeoCoord(t, validGeoCoord, g)
}

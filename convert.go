package h3

import (
	"unsafe"

	"github.com/akhenakh/goh3/ch3"
	"modernc.org/libc"
)

type Convert struct {
	*libc.TLS
}

func NewConvert() *Convert {
	return &Convert{TLS: libc.NewTLS()}
}

func (c *Convert) FromGeo(geo GeoCoord, res int) H3Index {
	cgeo := ch3.TGeoCoord{
		Flat: deg2rad * geo.Latitude,
		Flon: deg2rad * geo.Longitude,
	}

	return H3Index(ch3.XgeoToH3(c.TLS, uintptr(unsafe.Pointer(&cgeo)), int32(res)))
}

func (c *Convert) ToGeo(h H3Index) GeoCoord {
	cg := ch3.TGeoCoord{}
	ch3.Xh3ToGeo(c.TLS, ch3.TH3Index(h), uintptr(unsafe.Pointer(&cg)))
	g := GeoCoord{}
	g.Latitude = rad2deg * float64(cg.Flat)
	g.Longitude = rad2deg * float64(cg.Flon)

	return g
}

func (c *Convert) Close() {
	c.TLS.Close()
}

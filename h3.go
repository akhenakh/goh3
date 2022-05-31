package h3

import (
	"math"
	"unsafe"

	"github.com/akhenakh/goh3/ch3"
	"modernc.org/libc"
)

var (
	deg2rad = math.Pi / 180.0
	rad2deg = 180.0 / math.Pi
)

type H3Index ch3.TH3Index

type GeoCoord struct {
	Latitude, Longitude float64
}

type GeoPolygon struct {
	Geofence []GeoCoord

	Holes [][]GeoCoord
}

func FromGeo(geo GeoCoord, res int) H3Index {
	tls := libc.NewTLS()

	cgeo := ch3.TGeoCoord{
		Flat: deg2rad * geo.Latitude,
		Flon: deg2rad * geo.Longitude,
	}

	return H3Index(ch3.XgeoToH3(tls, uintptr(unsafe.Pointer(&cgeo)), int32(res)))
}

func ToGeo(h H3Index) GeoCoord {
	tls := libc.NewTLS()

	cg := ch3.TGeoCoord{}
	ch3.Xh3ToGeo(tls, ch3.TH3Index(h), uintptr(unsafe.Pointer(&cg)))
	g := GeoCoord{}
	g.Latitude = rad2deg * float64(cg.Flat)
	g.Longitude = rad2deg * float64(cg.Flon)

	return g
}

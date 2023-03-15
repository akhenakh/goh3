package h3

import (
	"unsafe"

	"github.com/akhenakh/goh3/ch3"
	"modernc.org/libc"
)

type Batch struct {
	tls *libc.TLS
}

func NewBatch() *Batch {
	return &Batch{tls: libc.NewTLS()}
}

func (b *Batch) LatLngToCell(latLng LatLng, resolution int) Cell {
	var i ch3.TH3Index

	cll := ch3.TLatLng{Flat: deg2rad * latLng.Lat, Flng: deg2rad * latLng.Lng}

	ch3.XlatLngToCell(b.tls, uintptr(unsafe.Pointer(&cll)), int32(resolution), uintptr(unsafe.Pointer(&i)))

	return Cell(i)
}

// CellToLatLng returns the geographic centerpoint of a Cell.
func (b *Batch) CellToLatLng(c Cell) LatLng {
	var g ch3.TLatLng

	ch3.XcellToLatLng(b.tls, ch3.TH3Index(c), uintptr(unsafe.Pointer(&g)))

	return LatLng{rad2deg * g.Flat, rad2deg * g.Flng}
}

func (b *Batch) Close() {
	b.tls.Close()
}

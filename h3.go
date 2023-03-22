package h3

import (
	"math"
	"strconv"
	"unsafe"

	"github.com/akhenakh/goh3/ch3"
	"modernc.org/libc"
)

const (
	numCellEdges = 6
	numEdgeCells = 2

	// MaxCellBndryVerts is the maximum number of vertices that can be used
	// to represent the shape of a cell.
	MaxCellBndryVerts = ch3.DMAX_CELL_BNDRY_VERTS
)

var (
	deg2rad = math.Pi / 180.0
	rad2deg = 180.0 / math.Pi
)

type DirectedEdge int64

type H3Index ch3.TH3Index

// Cell is an Index that identifies a single hexagon cell at a resolution.
type Cell uint64

// CellBoundary is a slice of LatLng.  Note, len(CellBoundary) will never
// exceed MaxCellBndryVerts.
type CellBoundary []LatLng

// GeoLoop is a slice of LatLng points that make up a loop.
type GeoLoop []LatLng

// LatLng is a struct for geographic coordinates in degrees.
type LatLng struct {
	Lat, Lng float64
}

type OptionsFunc func(*Options)

type Options struct {
	TLS *libc.TLS
}

// WithMultipleFences enable multi fences in responses
func WithTLS(tls *libc.TLS) OptionsFunc {
	return func(o *Options) {
		o.TLS = tls
	}
}

func tlsFromOption(opts []OptionsFunc) *libc.TLS {
	var lopts Options
	for _, opt := range opts {
		opt(&lopts)
	}

	return lopts.TLS
}

func NewLatLng(lat, lng float64) LatLng {
	return LatLng{lat, lng}
}

// LatLngToCell returns the Cell at resolution for a geographic coordinate.
func LatLngToCell(latLng LatLng, resolution int, opts ...OptionsFunc) Cell {
	var i ch3.TH3Index

	tls := tlsFromOption(opts)
	if tls == nil {
		tls = libc.NewTLS()
		defer tls.Close()
	}

	cll := ch3.TLatLng{Flat: deg2rad * latLng.Lat, Flng: deg2rad * latLng.Lng}

	ch3.XlatLngToCell(tls, uintptr(unsafe.Pointer(&cll)), int32(resolution), uintptr(unsafe.Pointer(&i)))

	return Cell(i)
}

func (c Cell) Resolution(opts ...OptionsFunc) int {
	tls := tlsFromOption(opts)
	if tls == nil {
		tls = libc.NewTLS()
		defer tls.Close()
	}

	return int(ch3.XgetResolution(tls, ch3.TH3Index(c)))
}

// CellToLatLng returns the geographic centerpoint of a Cell.
func CellToLatLng(c Cell, opts ...OptionsFunc) LatLng {
	tls := tlsFromOption(opts)
	if tls == nil {
		tls = libc.NewTLS()
		defer tls.Close()
	}

	var g ch3.TLatLng

	ch3.XcellToLatLng(tls, ch3.TH3Index(c), uintptr(unsafe.Pointer(&g)))

	return LatLng{rad2deg * g.Flat, rad2deg * g.Flng}
}

// LatLng returns the Cell at resolution for a geographic coordinate.
func (c Cell) LatLng(opts ...OptionsFunc) LatLng {
	return CellToLatLng(c, opts...)
}

func (c Cell) String() string {
	return strconv.FormatUint(uint64(c), 16)
}

// IsValid returns if a Cell is a valid cell (hexagon or pentagon).
func (c Cell) IsValid(opts ...OptionsFunc) bool {
	tls := tlsFromOption(opts)
	if tls == nil {
		tls = libc.NewTLS()
		defer tls.Close()
	}

	return c != 0 && ch3.XisValidCell(tls, uint64(c)) == 1
}

// Parent returns the parent or grandparent Cell of this Cell.
func (c Cell) Parent(resolution int, opts ...OptionsFunc) Cell {
	var out H3Index

	tls := tlsFromOption(opts)
	if tls == nil {
		tls = libc.NewTLS()
		defer tls.Close()
	}

	ch3.XcellToParent(tls, uint64(c), int32(resolution), uintptr(unsafe.Pointer(&out)))

	return Cell(out)
}

// Parent returns the parent or grandparent Cell of this Cell.
func (c Cell) ImmediateParent(opts ...OptionsFunc) Cell {
	return c.Parent(c.Resolution() - 1)
}

// Children returns the children or grandchildren cells of this Cell.
func (c Cell) Children(resolution int, opts ...OptionsFunc) []Cell {
	var outsz int

	tls := tlsFromOption(opts)
	if tls == nil {
		tls = libc.NewTLS()
		defer tls.Close()
	}

	ch3.XcellToChildrenSize(tls, uint64(c), int32(resolution), uintptr(unsafe.Pointer(&outsz)))
	out := make([]H3Index, outsz)

	ch3.XcellToChildren(tls, uint64(c), int32(resolution), uintptr(unsafe.Pointer(&out[0])))

	return cellsFromC(out, false, false)
}

func cellsFromC(chs []H3Index, prune, refit bool) []Cell {
	// OPT: This could be more efficient if we unsafely cast the C array to a
	// []H3Index.
	out := make([]Cell, 0, len(chs))

	for i := range chs {
		if prune && chs[i] <= 0 {
			continue
		}

		out = append(out, Cell(chs[i]))
	}

	if refit {
		// Some algorithms require a maximum sized array, but only use a subset
		// of the memory.  refit sizes the slice to the last non-empty element.
		for i := len(out) - 1; i >= 0; i-- {
			if out[i] == 0 {
				out = out[:i]
			}
		}
	}

	return out
}

func edgesFromC(chs []H3Index) []DirectedEdge {
	out := make([]DirectedEdge, 0, len(chs))

	for i := range chs {
		if chs[i] <= 0 {
			continue
		}

		out = append(out, DirectedEdge(chs[i]))
	}

	return out
}

func cellBndryFromC(cb *ch3.TCellBoundary) CellBoundary {
	g := make(CellBoundary, 0, MaxCellBndryVerts)
	for i := int32(0); i < cb.FnumVerts; i++ {
		g = append(g, latLngFromC(cb.Fverts[i]))
	}

	return g
}

func latLngFromC(cg ch3.TLatLng) LatLng {
	g := LatLng{}
	g.Lat = rad2deg * float64(cg.Flat)
	g.Lng = rad2deg * float64(cg.Flng)

	return g
}

// ImmediateChildren returns the children or grandchildren cells of this Cell.
func (c Cell) ImmediateChildren(opts ...OptionsFunc) []Cell {
	return c.Children(c.Resolution()+1, opts...)
}

// IsPentagon returns true if this is a pentagon.
func (c Cell) IsPentagon(opts ...OptionsFunc) bool {
	tls := tlsFromOption(opts)
	if tls == nil {
		tls = libc.NewTLS()
		defer tls.Close()
	}

	return ch3.XisPentagon(tls, uint64(c)) == 1
}

// IsNeighbor returns true if this Cell is a neighbor of the other Cell.
func (c Cell) IsNeighbor(other Cell, opts ...OptionsFunc) bool {
	var out int

	tls := tlsFromOption(opts)
	if tls == nil {
		tls = libc.NewTLS()
		defer tls.Close()
	}

	ch3.XareNeighborCells(tls, uint64(c), uint64(other), uintptr(unsafe.Pointer(&out)))
	return out == 1
}

// DirectedEdge returns a DirectedEdge from this Cell to other.
func (c Cell) DirectedEdge(other Cell, opts ...OptionsFunc) DirectedEdge {
	var out H3Index

	tls := tlsFromOption(opts)
	if tls == nil {
		tls = libc.NewTLS()
		defer tls.Close()
	}

	ch3.XcellsToDirectedEdge(tls, uint64(c), uint64(other), uintptr(unsafe.Pointer(&out)))
	return DirectedEdge(out)
}

// DirectedEdges returns 6 directed edges with h as the origin.
func (c Cell) DirectedEdges(opts ...OptionsFunc) []DirectedEdge {
	out := make([]H3Index, numCellEdges) // always 6 directed edges

	tls := tlsFromOption(opts)
	if tls == nil {
		tls = libc.NewTLS()
		defer tls.Close()
	}

	ch3.XoriginToDirectedEdges(tls, uint64(c), uintptr(unsafe.Pointer(&out[0])))
	return edgesFromC(out)
}

// ChildPosToCell returns the child cell at a given position within an ordered list of all
// children at the specified resolution res.
func (c Cell) ChildPosToCell(pos, res int, opts ...OptionsFunc) Cell {
	var out H3Index

	tls := tlsFromOption(opts)
	if tls == nil {
		tls = libc.NewTLS()
		defer tls.Close()
	}

	ch3.XchildPosToCell(tls, int64(pos), uint64(c), int32(res), uintptr(unsafe.Pointer(&out)))

	return Cell(out)
}

// ChildPos returns the position of the cell within an ordered list of all children of the cell's parent
// at the specified resolution res.
func (c Cell) ChildPos(res int, opts ...OptionsFunc) int {
	var out int

	tls := tlsFromOption(opts)
	if tls == nil {
		tls = libc.NewTLS()
		defer tls.Close()
	}

	ch3.XcellToChildPos(tls, uint64(c), int32(res), uintptr(unsafe.Pointer(&out)))

	return out
}

func (e DirectedEdge) IsValid(opts ...OptionsFunc) bool {
	tls := tlsFromOption(opts)
	if tls == nil {
		tls = libc.NewTLS()
		defer tls.Close()
	}

	return ch3.XisValidDirectedEdge(tls, uint64(e)) == 1
}

// Origin returns the origin cell of this directed edge.
func (e DirectedEdge) Origin(opts ...OptionsFunc) Cell {
	var out H3Index

	tls := tlsFromOption(opts)
	if tls == nil {
		tls = libc.NewTLS()
		defer tls.Close()
	}

	ch3.XgetDirectedEdgeOrigin(tls, uint64(e), uintptr(unsafe.Pointer(&out)))
	return Cell(out)
}

// Destination returns the destination cell of this directed edge.
func (e DirectedEdge) Destination(opts ...OptionsFunc) Cell {
	var out H3Index

	tls := tlsFromOption(opts)
	if tls == nil {
		tls = libc.NewTLS()
		defer tls.Close()
	}

	ch3.XgetDirectedEdgeDestination(tls, uint64(e), uintptr(unsafe.Pointer(&out)))
	return Cell(out)
}

// Cells returns the origin and destination cells in that order.
func (e DirectedEdge) Cells(opts ...OptionsFunc) []Cell {
	out := make([]H3Index, numEdgeCells)

	tls := tlsFromOption(opts)
	if tls == nil {
		tls = libc.NewTLS()
		defer tls.Close()
	}

	ch3.XdirectedEdgeToCells(tls, uint64(e), uintptr(unsafe.Pointer(&out[0])))

	return cellsFromC(out, false, false)
}

// Boundary provides the coordinates of the boundary of the directed edge. Note,
// the type returned is CellBoundary, but the coordinates will be from the
// center of the origin to the center of the destination. There may be more than
// 2 coordinates to account for crossing faces.
func (e DirectedEdge) Boundary(opts ...OptionsFunc) CellBoundary {
	var out ch3.TCellBoundary

	tls := tlsFromOption(opts)
	if tls == nil {
		tls = libc.NewTLS()
		defer tls.Close()
	}

	ch3.XdirectedEdgeToBoundary(tls, uint64(e), uintptr(unsafe.Pointer(&out)))

	return cellBndryFromC(&out)
}

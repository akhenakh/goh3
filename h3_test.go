package h3

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

const eps = 1e-4

const (
	// validH3Index resolution 5
	validCell    = Cell(0x850dab63fffffff)
	pentagonCell = Cell(0x821c07fffffffff)
)

var (
	validDiskDist3_1 = [][]Cell{
		{
			validCell,
		},
		{
			0x850dab73fffffff,
			0x850dab7bfffffff,
			0x850dab6bfffffff,
			0x850dab6ffffffff,
			0x850dab67fffffff,
			0x850dab77fffffff,
		},
		{
			0x850dab0bfffffff,
			0x850dab47fffffff,
			0x850dab4ffffffff,
			0x850d8cb7fffffff,
			0x850d8ca7fffffff,
			0x850d8dd3fffffff,
			0x850d8dd7fffffff,
			0x850d8d9bfffffff,
			0x850d8d93fffffff,
			0x850dab2bfffffff,
			0x850dab3bfffffff,
			0x850dab0ffffffff,
		},
	}

	validLatLng1 = LatLng{
		Lat: 67.1509268640,
		Lng: -168.3908885810,
	}
)

func almostEqual(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) {
	assert.InEpsilon(t, expected, actual, eps, msgAndArgs...)
}

func assertEqual[T comparable](t *testing.T, expected, actual T, msgAndArgs ...interface{}) {
	t.Helper()

	if expected != actual {
		var (
			expStr, actStr string

			e, a interface{} = expected, actual
		)

		switch e.(type) {
		case Cell:
			expStr = fmt.Sprintf("%s (res=%d)", e.(Cell), e.(Cell).Resolution())
			actStr = fmt.Sprintf("%s (res=%d)", a.(Cell), a.(Cell).Resolution())
		default:
			expStr = fmt.Sprintf("%v", e)
			actStr = fmt.Sprintf("%v", a)
		}
		t.Errorf("%v != %v", expStr, actStr)
		logMsgAndArgs(t, msgAndArgs...)
	}
}

func logMsgAndArgs(t *testing.T, msgAndArgs ...interface{}) {
	t.Helper()

	if len(msgAndArgs) > 0 {
		t.Logf(msgAndArgs[0].(string), msgAndArgs[1:]...)
	}
}

func assertCellIn(t *testing.T, needle Cell, haystack []Cell) {
	t.Helper()

	var found bool
	for _, h := range haystack {
		found = needle == h
		if found {
			break
		}
	}

	if !found {
		t.Errorf("%v not found in %+v", needle, haystack)
	}
}

func equalEps(expected, actual float64) bool {
	return math.Abs(expected-actual) < eps
}

func assertEqualEps(t *testing.T, expected, actual float64, msgAndArgs ...interface{}) {
	t.Helper()

	if !equalEps(expected, actual) {
		t.Errorf("%0.4f > %v (%0.4f - %0.4f)", math.Abs(expected-actual), eps, expected, actual)
		logMsgAndArgs(t, msgAndArgs...)
	}
}

func assertEqualLatLng(t *testing.T, expected, actual LatLng) {
	t.Helper()
	assertEqualEps(t, expected.Lat, actual.Lat, "latitude mismatch")
	assertEqualEps(t, expected.Lng, actual.Lng, "longitude mismatch")
}

func TestLatLngToCell(t *testing.T) {
	t.Parallel()
	c := LatLngToCell(validLatLng1, 5)
	assertEqual(t, validCell, c)
}

func TestCellToLatLng(t *testing.T) {
	t.Parallel()
	g := CellToLatLng(validCell)
	assertEqualLatLng(t, validLatLng1, g)
}

func TestIsValid(t *testing.T) {
	t.Parallel()
	assert.True(t, validCell.IsValid())
	assert.False(t, Cell(0).IsValid())
}

func TestParent(t *testing.T) {
	t.Parallel()
	// get the index's parent by requesting that index's resolution+1
	parent := validCell.ImmediateParent()

	// get the children at the resolution of the original index
	children := parent.ImmediateChildren()

	assertCellIn(t, validCell, children)
}

func TestIsPentagon(t *testing.T) {
	t.Parallel()
	assert.False(t, validCell.IsPentagon())
	assert.True(t, pentagonCell.IsPentagon())
}

func TestDirectedEdge(t *testing.T) {
	t.Parallel()

	origin := validDiskDist3_1[1][0]
	destination := origin.DirectedEdges()[0].Destination()
	edge := origin.DirectedEdge(destination)

	t.Run("is valid", func(t *testing.T) {
		t.Parallel()
		assert.True(t, edge.IsValid())
		assert.False(t, DirectedEdge(validCell).IsValid())
	})

	t.Run("get origin/destination from edge", func(t *testing.T) {
		t.Parallel()
		assertEqual(t, origin, edge.Origin())
		assertEqual(t, destination, edge.Destination())

		// shadow origin/destination
		cells := edge.Cells()
		origin, destination := cells[0], cells[1]
		assertEqual(t, origin, edge.Origin())
		assertEqual(t, destination, edge.Destination())
	})

	t.Run("get edges from hexagon", func(t *testing.T) {
		t.Parallel()
		edges := validCell.DirectedEdges()
		assertEqual(t, 6, len(edges), "hexagon has 6 edges")
	})

	t.Run("get edges from pentagon", func(t *testing.T) {
		t.Parallel()
		edges := pentagonCell.DirectedEdges()
		assertEqual(t, 5, len(edges), "pentagon has 5 edges")
	})

	t.Run("get boundary from edge", func(t *testing.T) {
		t.Parallel()
		gb := edge.Boundary()
		assertEqual(t, 2, len(gb), "edge has 2 boundary cells")
	})
}

func TestIsNeighbor(t *testing.T) {
	t.Parallel()
	assert.False(t, validCell.IsNeighbor(pentagonCell))
	assert.True(t, validCell.DirectedEdges()[0].Destination().IsNeighbor(validCell))
}

func TestChildPosToCell(t *testing.T) {
	t.Parallel()

	childrens := validCell.Children(6)

	assert.Equal(t, childrens[0], validCell.ChildPosToCell(0, 6))
}

func TestChildPos(t *testing.T) {
	t.Parallel()

	childrens := validCell.Children(7)

	assert.Equal(t, 32, childrens[32].ChildPos(5))
}

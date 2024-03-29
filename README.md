## goh3

This is an experiment h3 build using ccgo.

It's implementing two features that are not yet present in the current h3-go repo:
```go
func (c Cell) ChildPosToCell(pos, res int) Cell
func (c Cell) ChildPos(res int) int
```

WIP, use at your own risk

## Thread local storage

This lib can achieve speed close to the h3 Go using CGO library by batching, lowering pressure on allocating the thread local storage.

Each function is taking an optional `opts ...OptionsFunc`, you can pass your existing TLS: `WithTLS(tls)`.

```go
tls = libc.NewTLS()
defer tls.Close()

g := h3.CellToLatLng(validCell, h3.WithTLS(tls))
```

## Code Generation
```
  CC=/usr/bin/gcc ccgo  -pkgname ch3  -trace-translation-units -export-externs X -export-defines D -export-fields F -export-structs S -export-typedefs T -Isrc/h3lib/include  -I../src/h3lib/include ../src/h3lib/lib/*.c
```

## Patch h3 c sources to build on Linux (not needed on OSX)

Patch the original source to avoid builtin isfinite isssue (C tests are passing).
```C
bool isXfinite(double f) { return !isnan(f - f); }
```

Replace occurence of isfinite() with isXfinite() eg:
```C
/**
 * Encodes a coordinate on the sphere to the H3 index of the containing cell at
 * the specified resolution.
 *
 * Returns 0 on invalid input.
 *
 * @param g The spherical coordinates to encode.
 * @param res The desired H3 resolution for the encoding.
 * @return The encoded H3Index (or H3_NULL on failure).
 */
H3Index H3_EXPORT(geoToH3)(const GeoCoord* g, int res) {
    if (res < 0 || res > MAX_H3_RES) {
        return H3_NULL;
    }
    if (!isXFinite(g->lat) || !isXFinite(g->lon)) {
        return H3_NULL;
    }

    FaceIJK fijk;
    _geoToFaceIjk(g, res, &fijk);
    return _faceIjkToH3(&fijk, res);
}
```

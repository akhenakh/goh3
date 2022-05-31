## goh3

This is an experiment h3 build using ccgo.

WIP, use at your own risk

## Generation
```
  CC=/usr/bin/gcc ccgo  -pkgname ch3  -trace-translation-units -export-externs X -export-defines D -export-fields F -export-structs S -export-typedefs T -Isrc/h3lib/include  -I../src/h3lib/include ../src/h3lib/lib/*.c
```

## Patch h3 c sources

path the original source to avoid builtin isfinite isssue (tests are passing).
```
bool isXFinite(double f) { return !isnan(f - f); }

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
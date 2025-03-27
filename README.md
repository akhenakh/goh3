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
CC=/usr/bin/gcc ccgo  -Isrc/h3lib/include  -I../src/h3lib/include ../src/h3lib/lib/*.c    --package-name=ch3 --prefix-enumerator=_ --prefix-external=X --prefix-field=F --prefix-macro=D --prefix-static-internal=_ --prefix-static-none=_ --prefix-tagged-enum=_ --prefix-tagged-struct=T --prefix-tagged-union=T --prefix-typename=T --prefix-undefined=_
```


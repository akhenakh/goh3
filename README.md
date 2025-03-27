## goh3

This is an experimental Go native [h3](https://www.uber.com/en-CA/blog/h3/) build using ccgo.
It proves itself to be very useful since it does not require CGO.

It used in https://github.com/akhenakh/geo-bento.

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
tar xf v4.2.1.tar.gz
cd h3-4.2.1
cmake -B build
cd build
make -j 20
cp -Rp src/ ..
cp build/src/h3lib/include/h3api.h ../src/h3lib/include
cd ../ch3
CC=/usr/bin/gcc ccgo  -Isrc/h3lib/include  -I../src/h3lib/include ../src/h3lib/lib/*.c    --package-name=ch3 --prefix-enumerator=_ --prefix-external=X --prefix-field=F --prefix-macro=D --prefix-static-internal=_ --prefix-static-none=_ --prefix-tagged-enum=_ --prefix-tagged-struct=T --prefix-tagged-union=T --prefix-typename=T --prefix-undefined=_
```


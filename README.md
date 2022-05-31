## goh3

This is an experiment h3 build using ccgo.

WIP, use at your own risk

## Generation
```
  CC=/usr/bin/gcc ccgo  -pkgname ch3  -trace-translation-units -export-externs X -export-defines D -export-fields F -export-structs S -export-typedefs T -Isrc/h3lib/include  -I../src/h3lib/include ../src/h3lib/lib/*.c
```
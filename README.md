# Go Linters

Custom linters for the Go Programming Language.

## makeslicecap

`makeslicecap` helps avoid the common mistake of making a slice with an
initialized length rather than a capacity (e.g. calling `make([]int, 5)`
instead of `make([]int, 0, 5)`).

If you actually mean to call `make([]int, 5)` just spend the extra blink of
an eye to type the long-hand: `make([]int, 5, 5)`.

## errinspect

This is not a replacement for
[kisielk/errcheck](https://github.com/kisielk/errcheck).

`errinspect` helps avoid subtly ignoring errors. It is easy and legal to the
compiler to only use `err` to check for `err != nil` and no where else.

The following code would produce a linter error from `errinspect`:

```go
if err != nil {
  return 0
}
```

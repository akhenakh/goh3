package ch3

import (
	"math"

	"modernc.org/libc"
)

func Xfmin(tls *libc.TLS, a, b float64) float64 {
	return math.Min(a, b)
}

func Xfmax(tls *libc.TLS, a, b float64) float64 {
	return math.Max(a, b)
}

func Xlroundl(tls *libc.TLS, a float64) float64 {
	return math.Round(a)
}

//go:build go1.27
package main
type Box struct{}
func (Box) Pair[T any](v T) T{return v}

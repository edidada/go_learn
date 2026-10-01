package main
import "io"
var _ = struct{ io.Reader }.Read
func main() {}

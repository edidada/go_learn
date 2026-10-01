package main
type A interface{ M() }; type B interface{ M() }; type C interface{ A; B }
func main() {}

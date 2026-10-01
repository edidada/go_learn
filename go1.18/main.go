package main
func identity[T any](v T) T { return v }
func main(){ _=identity(18) }

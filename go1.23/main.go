package main
func seq(yield func(int) bool){ for i:=0;i<2;i++ { if !yield(i){return} } }
func main(){ for range seq {} }

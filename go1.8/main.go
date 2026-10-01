package main

import "fmt"

type jsonFoo struct {
	X int `json:"foo"`
}
type jsonBar struct {
	X int `json:"bar"`
}

func convertIgnoringTags(v jsonBar) jsonFoo { return jsonFoo(v) }

func main() { fmt.Println(convertIgnoringTags(jsonBar{X: 8})) }

package main

import "unsafe"
const (
	a = "abcbb"
	b = len(a)
	c = unsafe.Sizeof(a)
)

func main(){
	println(a, b, c)
}
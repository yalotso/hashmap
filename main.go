package main

import "fmt"

func main() {
	hm := NewHashMap()
	hm.Add(12, "Hello World")
	fmt.Println(hm.Get(12))
}

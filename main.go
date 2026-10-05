package main

import "fmt"

func main() {

	var a int8 = 127
	fmt.Println(a)
	var b int64 = 127
	fmt.Println(b)
	var c uint8 = 45
	fmt.Println(c)
	var d int8 = 45
	fmt.Println(d)
	d = -45
	fmt.Println(d)

	var dd = uint8(b)

	fmt.Println(dd)
}
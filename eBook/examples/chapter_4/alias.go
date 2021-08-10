package main

import fm "fmt" // alias

func Thread(a,b,c int, d ... string) (e int ){
	e = a + b +c
	fm.Println(d[0])
	return
}

var wlk, lk, tm = 1, 2, "xkn,"

func main() {
	fm.Println("hello, world")
	d := Thread(1, 2, 3, "hello world")
	fm.Println(d)
	fm.Println(wlk, lk, tm)
}

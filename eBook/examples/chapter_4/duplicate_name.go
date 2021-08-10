package main


import (
	fm "fmt"
	"os"
) // alias

const (
	M, TS, S = 1, 2, 3
	MQ = iota
	REDISSA
)
func Thread(a,b,c int, d ... string) (e int ){
	e = a + b +c
	var t [5]int= [5]int{}
	const Pi = len(t)
	fm.Println(d[0])
	return
}


var amsn int

func main(){
	Thread(1,2,3, "superman")
	fm.Println(M, TS, S)
	fm.Println(MQ)
	fm.Println(REDISSA)
	fm.Println(amsn)

}

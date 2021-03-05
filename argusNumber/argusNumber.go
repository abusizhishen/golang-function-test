package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println()
	run7()
	run8()
	run6()


}

func number6(a1,a2,a3,a4,a5,a6 int)  {

}

func number7(a1,a2,a3,a4,a5,a6,a7 int)  {

}

func number8(a1,a2,a3,a4,a5,a6,a7,a8 int)  {

}

func run6()  {
	var t = time.Now()
	for i:=0;i<1e8;i++{
		a1,a2,a3,a4,a5,a6 := 1,2,3,4,5,6
		number6(a1,a2,a3,a4,a5,a6)
	}

	fmt.Println("number6",time.Since(t))
}

func run7()  {
	var t = time.Now()
	for i:=0;i<1e8;i++{
		a1,a2,a3,a4,a5,a6,a7 := 1,2,3,4,5,6,7
		number7(a1,a2,a3,a4,a5,a6,a7)
	}

	fmt.Println("number7",time.Since(t))
}

func run8()  {
	var t = time.Now()
	for i:=0;i<1e8;i++{
		a1,a2,a3,a4,a5,a6,a7,a8 := 1,2,3,4,5,6,7,8
		number8(a1,a2,a3,a4,a5,a6,a7,a8)
	}

	fmt.Println("number8",time.Since(t))
}
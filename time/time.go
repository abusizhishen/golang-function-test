package main

import (
	"fmt"
	"time"
)

func main(){
	var s = "2021-03-05"
	t,err := time.Parse("2006-01-02",s)
	fmt.Println(t.String(),err)
}

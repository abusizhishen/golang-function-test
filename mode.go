package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println(check())
}


var mode string = "abb"

func check() bool {
	var str = "北京 南京 南京"
	var arr = strings.Split(str, " ")

	if len(arr) != len(mode){
		fmt.Println("length not equal")

		return false
	}

	var symbolValueMap = map[uint8]string{}
	var valueSymbolMap = map[string]uint8{}

	for i:=0;i<len(mode);i++{
		if value,ok := symbolValueMap[mode[i]];!ok {
			if _,ok := valueSymbolMap[arr[i]];ok{
				return false
			}

			symbolValueMap[mode[i]] = arr[i]
			valueSymbolMap[arr[i]] = mode[i]
			continue
		}else{
			if arr[i] != value{
				return false
			}
		}
	}

	return true
}
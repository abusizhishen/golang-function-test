package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"io"
)

func main() {
 	ReadFromFile("/Users/whw/Downloads/用户模板.xls")
}

func ReadFromReader(reader io.Reader) {
	f, err := excelize.OpenReader(reader)
	if err != nil {
		return
	}

	list := f.GetSheetList()
	if len(list) == 0 {
		err = fmt.Errorf("empty excel")
		return
	}


}

func ReadFromFile(name string) {
	f, err := excelize.OpenFile(name)
	if err != nil {
		panic(err)
		return
	}

	list := f.GetSheetList()
	if len(list) == 0 {
		err = fmt.Errorf("empty excel")
		return
	}


	fmt.Println(list)
}
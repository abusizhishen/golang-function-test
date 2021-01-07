package main

import (
	"github.com/bitly/go-simplejson"
	"log"
)

func main() {

	sJson := simplejson.New()
	sJson.Set("result", "success")
	rs := sJson.MustMap()
	log.Printf("%v",rs)
}

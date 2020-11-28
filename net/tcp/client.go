package main

import (
	"log"
	"net"
)

func main() {
	tcp()
}

func tcp()  {
	var conn,err = net.Dial("tcp",":90")
	if err != nil{
		log.Println(err)
		return
	}

	var s = "hello golang"
	_,err = conn.Write([]byte(s))
	log.Println(err)
}
package main

import (
	"log"
	"net"
)

func main() {
	server()
}

func server()  {
	var addr,err = net.ResolveTCPAddr("tcp", ":90")
	if err != nil{
		panic(err)
	}

	conn,err := net.ListenTCP("tcp", addr)
	if err != nil{
		panic(err)
	}

	var size = 1024
	var byts = make([]byte,size)
	for{
		tcpConn,err := conn.AcceptTCP()
		if err != nil{
			panic(err)
		}


		for {
			n,err := tcpConn.Read(byts)
			if err != nil{
				log.Println("read err")
				break
			}
			for n == size{
				log.Print(string(byts))
				continue
			}

			log.Println(string(byts[:n]))
			break
		}
	}
}

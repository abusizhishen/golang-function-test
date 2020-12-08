package main

import (
	"log"
	"net"
	"os"
)


func main()  {
	l,err := net.Listen("tcp",":102")
	if err != nil{
		log.Println(err)
		os.Exit(0)
	}

	for {
		conn,err := l.Accept()
		if err != nil{
			log.Printf("连接异常:%v",err)
			continue
		}

		log.Printf("new conn addr: %v\n", conn.RemoteAddr())
		go handler(conn)
	}
}

func handler(conn net.Conn)  {
	var byt = make([]byte,1024)
	var n int
	var err error
	defer conn.Close()

	for {
		n,err = conn.Read(byt)
		if err != nil{
			log.Printf("read err on addr:%v, err:%v", conn.RemoteAddr(),err)
			break
		}

		log.Println("read length:", n)
		if n > 0{
			log.Println(string(byt[:n]))
			continue
		}
		log.Println("-----read end-----")
		break
	}
}
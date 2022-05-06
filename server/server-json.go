package main

import (
	"go-rpc/structs"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	coll1 := structs.NewCollege()

	sv := rpc.NewServer()

	sv.Register(coll1)

	listener, _ := net.Listen("tcp", "localhost:1727")

	for {

		conn, _ := listener.Accept()

		go sv.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

package main

import (
	"go-rpc/common"
	"io"
	"net/http"
	"net/rpc"
)

func main() {
	coll1 := common.NewCollege()

	rpc.Register(coll1)

	rpc.HandleHTTP()

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "RPC SV LIVE!")
	})

	http.ListenAndServe(":1626", nil)
}

package main

import (
	"fmt"
	"go-rpc/structs"
	"net/rpc/jsonrpc"
)

func main() {
	client, _ := jsonrpc.Dial("tcp", "localhost:1727")

	var st1 structs.Student

	if err := client.Call("College.Get", 1, &st1); err != nil {
		fmt.Println("Error: 1 College.Get()", err)
	} else {
		fmt.Printf("Success:1 Student %s found with id=1 \n", st1.FullName())
	}

	if err := client.Call("College.Add", structs.Student{
		Id:        1,
		FirstName: "Joni",
		LastName:  "Moncada",
	}, &st1); err != nil {
		fmt.Println("Error:2 College.Add()", err)
	} else {
		fmt.Printf("Success:2 Student %s created with id=1 \n", st1.FullName())
	}
}

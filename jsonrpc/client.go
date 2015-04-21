package main

import (
	"code.google.com/p/go.net/websocket"
	"net/rpc/jsonrpc"

	"log"
)

type Args struct {
	A, B int
}

func main() {
	origin := "http://localhost"
	url := "ws://localhost:7000/conn"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	client := jsonrpc.NewClient(ws)

	args := &Args{1, 2}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Result: %d+%d=%d\n", args.A, args.B, reply)
}

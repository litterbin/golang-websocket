package main

import (
	"code.google.com/p/go.net/websocket"
	//"github.com/garyburd/go-websocket/websocket"
	//"github.com/zhangpeihao/gowebsocket"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"

	"log"
)

type Args struct {
	A int
	B int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	log.Println(args.A, args.B)
	*reply = args.A * args.B
	log.Printf("%d", *reply)
	return nil
}

func main() {
	rpc.Register(new(Arith))

	http.Handle("/conn", websocket.Handler(serve))
	http.ListenAndServe("localhost:7000", nil)
}

func serve(ws *websocket.Conn) {
	jsonrpc.ServeConn(ws)
}

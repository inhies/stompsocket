package main

import (
	"fmt"
	websocket "github.com/inhies/gowebsocket"
	"github.com/jjeffery/stomp/server"
	"net"
	"net/http"
)

type Listener chan net.Conn

var l = make(Listener)

func main() {
	// Start the STOMP server
	go server.Serve(l)

	// Websocket location
	http.HandleFunc("/ws", STOMPserver)

	// Chat files are in /chat/ but there's no need to have to use that URL
	// so we will serve them from /
	http.Handle("/", http.StripPrefix("/",
		http.FileServer(http.Dir("chat/"))))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func STOMPserver(w http.ResponseWriter, req *http.Request) {
	// Create a net.Conn from the websocket
	conn, err := websocket.NewConn(w, req, nil, 4096, 4096)
	if err != nil {
		fmt.Println("error getting websocket conn:", err)
		return
	}

	// Send this connection to our STOMP server
	l <- conn
}

func (l Listener) Accept() (c net.Conn, err error) {
	return <-l, nil
}

func (l Listener) Close() (err error) {
	return nil
}

func (l Listener) Addr() (a net.Addr) {
	return nil
}

Chat via STOMP over WebSocket
=============================

A simple chat server and client example using Go and JavaScript.

Usage
-----

Fetch and install the package:

`go get github.com/inhies/stompsocket`

Then just run `stompsocket` which will start a webserver on port 8080 and point 
your browser to http://localhost:8080 and press connect to connect to the 
service. 

Credits
-------

The demo program came from the source for
[stomp.js](https://github.com/jmesnil/stomp-websocket)

The websocket code is mainly
[garyburd's](http://godoc.org/github.com/garyburd/go-websocket/websocket) with a
wrapper to create a net.Conn from
[zhangpeihao](http://godoc.org/github.com/zhangpeihao/gowebsocket). 

The STOMP code is from
[jjeffery](http://godoc.org/github.com/jjeffery/stomp/server).

Each project is licensed under their respective licenses.

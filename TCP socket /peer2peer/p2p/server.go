package p2p

import (
	"net"
	"fmt"
	"encoding/json"
	// "bytes"
	"log"
)

type Server struct {
	port int
	stream net.Listener
}

func (server *Server) SetPort(port int) {
	server.port = port
} 

func (server *Server) GetPort() int {
	return server.port
}

func (server *Server) Start() {
	randPort := PortDuration{49152,65535}
	var port int
	for {
		port = randPort.RandPort()
		stream, err := net.Listen("tcp",fmt.Sprintf(":%d", port))
		if err == nil {
			server.port = port
			server.stream = stream
			break
		} 
	}
	fmt.Println("Listening on port", server.port)
}

func (server *Server) Stop() {
	server.stream.Close()
	fmt.Println("Stop listening on port", server.port)
}

func (server *Server) Listener() {
	for {
		con, err := server.stream.Accept()
		if err != nil {
			return
		}else{
			go request(con)
		}
	}
}

type Json struct {
	Method string `json:"method"`
	Message string `json:"message"`
}

func request(con net.Conn) {
	defer fmt.Println("Socket is closed")
	defer con.Close()
	for {
		fmt.Println("wait")
		readBuf := make([]byte, 4096)
		len, err := con.Read(readBuf)
		if (err == nil) {
			var value map[string]interface{}
			err = json.Unmarshal(readBuf[:len], &value)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(readBuf))
			response(con, value)
			break
		}
		
	}		
}

func response(con net.Conn, value map[string]interface{}) {
	msgJson := map[string]interface{}{
		"method": value["method"],
		"message": "I'm very busy",
	}
	jsonData, err := json.Marshal(msgJson)
	checkErr(err)
    con.Write(jsonData)
	con.Close()
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}
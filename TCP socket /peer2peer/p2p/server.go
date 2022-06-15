package p2p

import (
	"net"
	"fmt"
	// "bytes"
	"github.com/golang/protobuf/proto"
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

func request(con net.Conn) {
	defer fmt.Println("Socket is closed")
	defer con.Close()
	for {
		fmt.Println("wait")
		readBuf := make([]byte, 4096)
		len, err := con.Read(readBuf)
		if (err == nil) {
			responseData := &Message{}
			err = proto.Unmarshal(readBuf[:len], responseData)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(readBuf))
			fmt.Println(responseData)
			response(con, responseData)
			break
		}
		
	}		
}

func response(con net.Conn, responseData *Message) {
	msgProtobuf := &Message{
		Method: responseData.GetMethod(),
		Message: "I'm very busy",
	}
	jsonData, err := proto.Marshal(msgProtobuf)
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
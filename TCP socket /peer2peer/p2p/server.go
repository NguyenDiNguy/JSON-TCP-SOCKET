package p2p

import (
	"net"
	"fmt"
	"bytes"
	"github.com/golang/protobuf/proto"
	"log"
	"encoding/binary"
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

func captureFullMessage(con net.Conn) (*Message, error) {
		bl := make([]byte, 4)
		_, err := con.Read(bl)
		if err != nil {
			return nil, err
		}
		rLen := binary.LittleEndian.Uint32(bl)
		mlen := rLen
		fmt.Println("len",rLen)
		var buf bytes.Buffer
		for rLen > 0 {
			
			readChilBuf := make([]byte, 100000000)
			len, err := con.Read(readChilBuf)
			fmt.Println(len, err)
			
			fmt.Println("len", rLen, "-", len)
			
			if (err == nil) {
				rLen-=uint32(len)
				buf.Write(readChilBuf[:len])
			}else{ break }
			
		}
		requestData := &Message{}
		err = proto.Unmarshal(buf.Next(int(mlen)), requestData)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		return requestData, nil
}

func request(con net.Conn) {
	defer fmt.Println("Socket is closed")
	defer con.Close()
	for {
		fmt.Println("wait")
		requestData := &Message{}
		requestData, err:= captureFullMessage(con)
		if err != nil {
			fmt.Println(err)
			return		
		}
		response(con, requestData)
	}		
}
func response(con net.Conn, requestData *Message) {
	msgProtobuf := &Message{
		Method: requestData.GetMethod(),
		Message: fmt.Sprintf("I'm very busy with %d", len(requestData.Message)),
	}
	jsonData, err := proto.Marshal(msgProtobuf)
	checkErr(err)
    con.Write(jsonData)
}





func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}
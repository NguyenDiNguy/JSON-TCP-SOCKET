package p2p

import (
	"fmt"
	"net"
	"github.com/golang/protobuf/proto"
	"../randtext"
	"encoding/binary"
)

type P2P struct {
	server *Server
}

func (p2p *P2P) Start() {
	p2p.server = new(Server)
	p2p.server.Start()
	go p2p.server.Listener()
}

func (p2p *P2P) Close() {
	p2p.server.Stop()
}

func (p2p *P2P) Send(port int, message string) {
	conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d",port))
	if err != nil {
		fmt.Println("Send message fail by", err)
		return
	}
	defer conn.Close()
	defer fmt.Println("Socket is closed")
	fmt.Println("Client ",conn.LocalAddr()," connected succesfull")
	for i:=0;i<3;i++ {
		sendRandomText(conn)
	}
} 

func sendRandomText(con net.Conn){
	msgProtobuf := &Message{
		Method: "Get",
		Message: randtext.RandStringBytesMaskImprSrcUnsafe(5050505),
	}
	jsonData, err := proto.Marshal(msgProtobuf)
	if (err != nil){ return }
	l := uint32(len(jsonData))
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf[0:], l)
	buf = append(buf, jsonData...)
    con.Write(buf)
	fmt.Println("wait")
	readBuf := make([]byte, 4096)
	len, err := con.Read(readBuf)
	if (err == nil) {
		response := &Message{}
		err = proto.Unmarshal(readBuf[:len], response)
		fmt.Println(string(readBuf))
	}
}

package p2p

import (
	"fmt"
	"net"
	"encoding/json"
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
	msgJson := map[string]interface{}{
		"method": "Get",
		"message": message,
	}
	jsonData, err := json.Marshal(msgJson)
	checkErr(err)
    conn.Write(jsonData)
	fmt.Println("wait")
	readBuf := make([]byte, 4096)
	len, err := conn.Read(readBuf)
	if (err == nil) {
		var value map[string]interface{}
		err = json.Unmarshal(readBuf[:len], &value)
		fmt.Println(string(readBuf))
	}
	conn.Close()
} 

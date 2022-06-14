package main

import (
    "fmt"
    "log"
    "net"
	"encoding/json"
	"sync"
	"bytes"
)
	
	func main()  {
		var wg sync.WaitGroup
		conn, err := net.Dial("tcp", "127.0.0.1:8080")
		checkErr(err)
		fmt.Println("Client ",conn.LocalAddr()," connected succesfull")
		defer conn.Close()
		defer fmt.Println("Socket is closed")
		go listen(conn, &wg) 
		wg.Add(1)
		msgJson := map[string]interface{}{
			"value": map[string]interface{}{
				"ip": "aaaaa",
				"port": 8080,
			},
		}
		jsonData, err := json.Marshal(msgJson)
    	conn.Write(jsonData)
		wg.Wait()
	}
	
	func checkErr(err error) {
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
	}

	type Data struct {
		Ip string `json:"ip"`
		Port int `json:"port"`
	}

	type Json struct {	
		Value []Data `json:"value"`
	}

	func listen(con net.Conn, wg *sync.WaitGroup) {
		readBuf := make([]byte, 4096)
		for {
			con.Read(readBuf)
			fmt.Println("wait");
			if string(readBuf[:5]) == "Close" {break}
			var value map[string]interface{}
			err := json.Unmarshal(bytes.Trim(readBuf[:len(readBuf)],"\x00"), &value)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(string(readBuf))
			con.Write([]byte("Close"))
		}
		wg.Done()
	}
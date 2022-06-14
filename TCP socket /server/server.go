package main

import (
    "fmt"
    "log"
    "net"
	"encoding/json"
	"bytes"
)

	func main()  {
		dstream, err := net.Listen("tcp", ":8080")
		fmt.Println("Listening on port", 8080)
		checkErr(err)
		defer dstream.Close()
		defer fmt.Println("Server down")
		for {
			con, err := dstream.Accept()
			checkErr(err)
			go handle(con)
		}
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
		Value Data `json:"value"`
	}

	func handle(con net.Conn) {
		defer fmt.Println("Socket is closed")
		defer con.Close()
		readBuf := make([]byte, 4096)
		for {
			_, err := con.Read(readBuf)
			fmt.Println("wait")
			if (err == nil) {
				if string(readBuf[:5]) == "Close" {
					con.Write([]byte("Close"))
					break
				}
				var value map[string]interface{}
				err = json.Unmarshal(bytes.Trim(readBuf[:len(readBuf)],"\x00"), &value)
				
				byteArray, err := json.Marshal(value)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(string(byteArray))
				if data, ok := value["value"].(map[string]interface{}); ok {
					data["port"] = 9999
					value["value"] = data
					byteArray, err := json.Marshal(value)
					if err != nil {
						fmt.Println(err)
						return
					}
					con.Write(byteArray)
				}		
			}
			
		}		
	}
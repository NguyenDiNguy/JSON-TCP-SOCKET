package main

import (
	"./p2p"
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)

var function = map[string](string){
	"1": "Khởi chạy server",
	"2": "Đóng server",
	"3": "Send message",
}

func main() {
	p2p := new(p2p.P2P)
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter action: ")
		for k,v := range function {
			fmt.Printf("\t%s. %s\n",k,v)
		} 
		text, _ := reader.ReadString('\n')
		option := strings.TrimSpace(text)
		switch option {
			case "1": 
				fmt.Println(function[option])
				p2p.Start()
			case "2":
				fmt.Println(function[option])
				p2p.Close()
			case "3": 
				fmt.Println(function[option])
				port, message := getMessage()
				p2p.Send(port, message)
			default:
				fmt.Println("Không có chức năng này")
		}
		
	}
}

func getMessage() (port int, message string) {
	reader := bufio.NewReader(os.Stdin)
	for { 
		fmt.Println("Sent to port: ")
		text, _ := reader.ReadString('\n')
		port, err := strconv.Atoi(strings.TrimSpace(text))
		if (err != nil) {
			fmt.Println("Vui lòng nhập lại")
			continue
		}
		fmt.Println("Message: ")
		text, _ = reader.ReadString('\n')
		return port, text
	}
}
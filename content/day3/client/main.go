package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const host = "localhost:9999"

func main() {
	reader := bufio.NewReader(os.Stdin)
	for{
		fmt.Printf("enter some text: ")
		data,err := reader.ReadString('\n')
		if err!= nil {
			fmt.Printf("encountered an error reading: %s\n",err.Error())
			continue
		}
		conn,err :=net.Dial("tcp",host)
		if err!=nil{
			fmt.Printf("error while connecting to server: %s\n",err.Error())
		
		}
		fmt.Fprint(conn,data)
		resp,err := bufio.NewReader(conn).ReadString('\n')
		if err!= nil {
			fmt.Printf("error while reading resp: %s",err.Error())
		}
		fmt.Printf("Recieved back: %s",resp)
		conn.Close()
	}
}
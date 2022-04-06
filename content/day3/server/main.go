package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"
)

const host = "localhost:9999"

func echoUppercase(conn net.Conn) ([]byte,error) {
	reader := bufio.NewReader(conn)

	data,err := reader.ReadString('\n')
	if err !=nil {
		
		return nil,errors.New("Something went wrong")
	}
	fmt.Printf("Received: %s",data)

	return []byte(strings.ToUpper(data)),nil
}

func main()  {
	listener, err:=net.Listen("tcp",host)
	if err !=nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Printf("listening on: %s\n",host)
	for {
		conn,err := listener.Accept()
		if err != nil {
			fmt.Printf("encountered an erro accepting connection: %s/n",err.Error())
			continue
		}
		upped,err:= echoUppercase(conn)
		if err!= nil {
			fmt.Printf("error reading data: %s\n",err.Error())
		}
		conn.Write(upped)
		conn.Close()
	}
}
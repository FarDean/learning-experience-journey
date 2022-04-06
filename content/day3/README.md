## Network programming

### Simple tcp server
building a tcp server that capitalizes the text received from client and sends it back.

1. create folder named server and a file in it named main.go
2. create a function that takes connection as argument(`net.Conn`) and return a `[]byte` and an `error`
```go
func echoCapitalized(conn net.Conn) ([]byte,error){}
```
3. define a Reader for the connection
```go
	reader := bufio.NewReader(conn)
```
4. read the received data
```go
data,err := reader.ReadString('\n')
```
5. handle the error
```go
	if err !=nil {
		return nil,errors.New("Something went wrong")
	}
```
6. print the received text on the server
```go
fmt.Printf("Received: %s",data)
```
7. return the uppercased text and a nil error
```go
return []byte(strings.ToUpper(data)),nil
```
8. create the main func
9. define a constant for host address
```go
const host = "localhost:9999"
```
10. create a new listener on the defined host
```go
listener,err := net.Listen("tcp",host)
```
11. handle the error
```go
	if err !=nil {
		panic(err)
	}
```
12. defer closing connection
```go 
defer listener.Close()
```
13. create a (infinite in this case) loop
```go
for{}
```
14. in the loop wait for the connections
```go
conn,err := listener.Accept()
```
15. handle the error
```go
		if err != nil {
			fmt.Printf("encountered an erro accepting connection: %s/n",err.Error())
			continue
		}
```
16. capitalize the received input using the function created earlier
```go
upped,err:= echoUppercase(conn)
```
19. handle the error
```go
		if err!= nil {
			fmt.Printf("error reading data: %s\n",err.Error())
		}
```
20. send the uppercased data to client
```go
conn.Write(upped)
```
21. close the connection
```go
conn.Close()
```

### Simple tcp client

1. create folder named client and a file in it named main.go
2. create the main func
3. define a Reader for stdin
```go
r:= bufio.NewReader(os.Stdin)
```
4. create a loop
5. get the user input from stdin
```go
data,err := r.ReadString('\n')
```
6. handle the error
```go
		if err!= nil {
			fmt.Printf("encountered an error reading: %s\n",err.Error())
			continue
		}
```
7. connect to the server
```go
conn,err :=net.Dial("tcp",host)
```
8. handle the error
```go
		if err!=nil{
			fmt.Printf("error while connecting to server: %s\n",err.Error())
		
		}
```
9. send the data to the server
```go
fmt.Fprint(conn,data)
```
10. get the response from the server and print it
```go
		resp,err := bufio.NewReader(conn).ReadString('\n')
		if err!= nil {
			fmt.Printf("error while reading resp: %s",err.Error())
		}
		fmt.Printf("Recieved back: %s",resp)
```
11. close the connection
```go
conn.Close()
```
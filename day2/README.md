## Simple http server
1. create a new folder and run:
```sh
$ go mod init app
```
2. create a file named `main.go`
3. import the `net/http` module
4. for a simple http server we need 3 things:
  - a handler function for a route
  - `http.HandleFunc` to use the handler function on certain route
  - `http.ListenAndServe` to serve the connection on a port

## Containerizing the server
1. create a `Dockerfile`
2. indicate the official `golang image` as the base image.
```Dockerfile
FROM golang:1.18
```
3. set the  `WORKDIR` to `/app` folder in the container
```Dockerfile
WORKDIR /app
```
4. inform Docker that the application listens for connection on a certain port
```Dockerfile
EXPOSE 8080
```
5. copy the project files into the `workdir` of the container
```Dockerfile
COPY main.go go.mod /app/
```
6. build the application inside the container
```Dockerfile
RUN go build -v -o /usr/local/bin/app .
```
7. execute the command to start the app inside the container
```Dockerfile
CMD ["app"]
```
8. build the image
```bash
docker build -t simple-http-go .
```
9. run the image
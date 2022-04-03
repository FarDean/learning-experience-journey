FROM golang:1.18
WORKDIR /app
EXPOSE 8080
COPY main.go go.mod /app/
RUN go build -v -o /usr/local/bin/app .
CMD [ "app" ]
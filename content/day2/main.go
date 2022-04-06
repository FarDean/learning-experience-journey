package main

import (
	"fmt"
	"log"
	"net/http"
)
const port = ":10000"

func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w,"Day two!")
}

func main()  {
	http.HandleFunc("/",handler)
	log.Fatal(http.ListenAndServe(port,nil))
}
package main

import (
	"fmt"
	"go-video-chat/endpoint"
	"log"
	"net/http"
)

func main() {
	fmt.Println("started")
	endpoint.NewEndpoint()
	log.Fatal(http.ListenAndServe(":8090", nil))
}

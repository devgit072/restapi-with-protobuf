package main

import (
	"bytes"
	hellopb "github.com/devraj/restapi-with-protobuf/protos"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"net/http"
)

/*
Steps:
1) Create proto request as Golang struct.
2) Convert that proto struct into bytes using proto.Marshal() method.
3) HTTP server will returns back bytes as response.
4) Convert that bytes into proto struct using method proto.Unmarshal
 */

// Content type can be set as application/x-binary or preferably application/x-protobuf

func main() {
	req := hellopb.HelloRequest{
		Request: "Hello Mars!",
		Greeter: "Elon Musk",
		Age:     45,
	}
	b, err := proto.Marshal(&req)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	// response, err := http.Post("http://localhost:8080/hello", "application/x-binary", bytes.NewReader(b))
	response, err := http.Post("http://localhost:8080/hello", "application/x-protobuf", bytes.NewReader(b))
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	b, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	helloResponse := &hellopb.HelloResponse{}
	if err := proto.Unmarshal(b, helloResponse); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	log.Printf("Response: \n%+v\n", helloResponse)
}

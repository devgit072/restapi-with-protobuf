package main

import (
	"fmt"
	hellopb "github.com/devraj/restapi-with-protobuf/protos"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", Hello).Methods("POST")
	router.HandleFunc("/ping", Ping).Methods("GET")
	s := &http.Server{
		Addr:              "localhost:8080",
		Handler:           router,
	}
	log.Println("Starting a new HTTP server on port 8080...")
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("Error: %+v", err)
	}
}

func Ping(resp http.ResponseWriter, req *http.Request) {
	if _, err := resp.Write([]byte("Pong!")); err != nil {
		log.Fatalf("Error: %+v", err)
	}
}

/*
Steps:
1) Read the request as bytes and convert that into proto using proto.Unmarshal()
2) Create a proto response and convert into bytes using proto.Marshal()
3) Send bytes back to the client.
 */
func Hello(resp http.ResponseWriter, req *http.Request) {
	d, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	helloReq := &hellopb.HelloRequest{}
	if err := proto.Unmarshal(d, helloReq); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	response := hellopb.HelloResponse{
		Response:  fmt.Sprintf("Response of req: %s", helloReq.Request),
		Responder: fmt.Sprintf("Martian, a frind of %s", helloReq.Greeter),
		Age:       20,
	}
	// Send back this response in serialized form.
	serializedData, err := proto.Marshal(&response)
	if err != nil {
		log.Fatalf("Error: %+v", err)
	}
	if _, err := resp.Write(serializedData); err != nil {
		log.Fatalf("Error: %+v", err)
	}
}

Example of how to implement http rest api in Go with protobuf data format. 
 
Protobuf is modern data format and it has proven much faster and has much smaller payload size than JSON. 
  
<b>Generate the proto code using:</b>  
```protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protos/*.proto```   

Workflow in http client: 

 Steps: 
 1) Create proto request as Golang struct.
 2) Convert that proto struct into bytes using proto.Marshal() method.
 3) HTTP server will returns back bytes as response.
 4) Convert that bytes into proto struct using method proto.Unmarshal
 
 Workflow in http server: 
 
 Steps:
 1) Read the request as bytes and convert that into proto using proto.Unmarshal()
 2) Create a proto response and convert into bytes using proto.Marshal()
 3) Send bytes back to the client.
 
 <b>Note:</b> 
 
 Content type can be set as application/x-binary or preferably application/x-protobuf
 
 

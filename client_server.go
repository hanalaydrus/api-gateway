/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

package main

import (
	"fmt"
	"time"
	"net/http"
	"log"
	"os"
	"io"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"

	pb "./helloworld"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc/reflection"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

var count = "0"
var prev_count = "-1"

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(in *pb.HelloRequest, stream pb.Greeter_SayHelloServer) error {
	stream.SendHeader(metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-stream"))
	for {
		if prev_count != count {
			helloReply := &pb.HelloReply{Message: count}
			if err := stream.Send(helloReply); err != nil {
				return err
			}
			log.Println("Sending")
			prev_count = count
		}
	}
	stream.SetTrailer(metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-stream"))
	return nil
}

func asClient() {
	////////// Client ////////////
	conn, erro := grpc.Dial(address, grpc.WithInsecure())
	if erro != nil {
		log.Fatalf("did not connect: %v", erro)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	stream, erro := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if erro != nil {
		log.Fatalf("could not greet: %v", erro)
	}
	for {
		helloReply, erro := stream.Recv()
		if erro == io.EOF {
			break
		}
		if erro != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", c, erro)
		}
		count = helloReply.Message
		log.Println("Car Count: ", count)
		time.Sleep(100 * time.Millisecond)
	}
}

func asServer() {
	//////////// Server //////////////
	port := 8080
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	wrappedServer := grpcweb.WrapServer(s, grpcweb.WithCorsForRegisteredEndpointsOnly(false))
	handler := func(resp http.ResponseWriter, req *http.Request) {
		wrappedServer.ServeHTTP(resp, req)
	}
	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(handler),
	}
	grpclog.Printf("Starting server. http port: %d", port)
	
	if err := httpServer.ListenAndServe(); err != nil {
		grpclog.Fatalf("failed starting http server: %v", err)
	}
}

func main() {
	go asClient()
	asServer()
}
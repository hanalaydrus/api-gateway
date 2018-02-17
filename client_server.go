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

package main

import (
	// "fmt"
	"time"
	// "net/http"
	"log"
	"io"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	// "google.golang.org/grpc/grpclog"
	// "google.golang.org/grpc/metadata"

	pbs "./semanticContract"
	// pbv "./volumeContract"
	// "github.com/improbable-eng/grpc-web/go/grpcweb"
	// "google.golang.org/grpc/reflection"
)

const (
	addressVolume = "localhost:50051"
	addressDensity = "localhost:50050"
	addressSemantic = "localhost:50049"
)

var count int32 = 0
var prev_count = -1

var timestamp = ""

var density_state = "Lancar"
var prev_density_state = "null"

var semantic_state = ""
var prev_semantic_state = "null"

// server is used to implement helloworld.GreeterServer.
// type server struct{}

// SayHello implements helloworld.GreeterServer
// func (s *server) SayHello(in *pb2.HelloRequest, stream pb2.Greeter_SayHelloServer) error {
// 	stream.SendHeader(metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-stream"))
// 	for {
// 		if ((prev_count != count) || (prev_density_state != density_state)) {
// 			helloReply := &pb2.HelloReply{Volume: count, Density: density_state}
// 			if err := stream.Send(helloReply); err != nil {
// 				return err
// 			}
// 			log.Println("Sending")
// 			prev_count = count
// 			prev_density_state = density_state
// 		}
// 	}
// 	stream.SetTrailer(metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-stream"))
// 	return nil
// }

// func asClientDensity() {
// 	var withBlock = grpc.WithBlock()
// 	conn, erro := grpc.Dial(addressDensity, grpc.WithInsecure(), withBlock)
// 	if erro != nil {
// 		log.Fatalf("did not connect: %v", erro)
// 	}
// 	defer conn.Close()
// 	c := pb.NewGreeterClient(conn)

// 	// Contact the server and print out its response.

// 	stream, erro := c.SayHello(context.Background(), &pb.HelloRequest{Id: 1})
// 	if erro != nil {
// 		log.Fatalf("could not greet: %v", erro)
// 	}
// 	for {
// 		helloReply, erro := stream.Recv()
// 		if erro == io.EOF {
// 			break
// 		}
// 		if erro != nil {
// 			log.Fatalf("%v.ListFeatures(_) = _, %v", c, erro)
// 		}
// 		timestamp = helloReply.Timestamp
// 		density_state = helloReply.Response
// 		log.Println("Timestamp: ", timestamp)
// 		log.Println("Density: ", density_state)
// 		time.Sleep(10 * time.Millisecond)
// 	}
// }

// func asClientVolume() {
// 	////////// Client ////////////
// 	var withBlock = grpc.WithBlock()
// 	conn, erro := grpc.Dial(addressVolume, grpc.WithInsecure(), withBlock)
// 	if erro != nil {
// 		log.Fatalf("did not connect: %v", erro)
// 	}
// 	defer conn.Close()
// 	c := pbv.NewGreeterClient(conn)

// 	// Contact the server and print out its response.
// 	stream, erro := c.SayHello(context.Background(), &pbv.HelloRequest{Id: 1})
// 	if erro != nil {
// 		log.Fatalf("could not greet: %v", erro)
// 	}
// 	for {
// 		helloReply, erro := stream.Recv()
// 		if erro == io.EOF {
// 			break
// 		}
// 		if erro != nil {
// 			log.Fatalf("%v.ListFeatures(_) = _, %v", c, erro)
// 		}
// 		count = helloReply.Volume
// 		log.Println("Car Count: ", count)
// 		time.Sleep(10 * time.Millisecond)
// 	}
// }

func asClientSemantic() {
	log.Println("tes")
	var withBlock = grpc.WithBlock()
	conn, erro := grpc.Dial(addressSemantic, grpc.WithInsecure(), withBlock)
	if erro != nil {
		log.Fatalf("did not connect: %v", erro)
	}
	defer conn.Close()
	c := pbs.NewGreeterClient(conn)

	// Contact the server and print out its response.

	stream, erro := c.SayHello(context.Background(), &pbs.HelloRequest{Id: 1})
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
		semantic_state = helloReply.Response
		log.Println("Semantic: ", semantic_state)
		time.Sleep(10 * time.Millisecond)
	}
}

// func asServer() {
// 	//////////// Server //////////////
// 	port := 8080
// 	s := grpc.NewServer()
// 	pb2.RegisterGreeterServer(s, &server{})
// 	// Register reflection service on gRPC server.
// 	reflection.Register(s)

// 	wrappedServer := grpcweb.WrapServer(s, grpcweb.WithCorsForRegisteredEndpointsOnly(false))
// 	handler := func(resp http.ResponseWriter, req *http.Request) {
// 		wrappedServer.ServeHTTP(resp, req)
// 	}
// 	httpServer := http.Server{
// 		Addr:    fmt.Sprintf(":%d", port),
// 		Handler: http.HandlerFunc(handler),
// 	}
// 	grpclog.Printf("Starting server. http port: %d", port)
	
// 	if err := httpServer.ListenAndServe(); err != nil {
// 		grpclog.Fatalf("failed starting http server: %v", err)
// 	}
// }

func main() {
	// go asClientVolume()
	// go asClientDensity()
	asClientSemantic()
	// asServer()
}
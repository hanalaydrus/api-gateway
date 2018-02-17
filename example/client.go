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
	"log"
	// "os"
	"io"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	// pb "google.golang.org/grpc/examples/helloworld/helloworld"
	pb "../helloworld"
)

const (
	address     = "localhost:8080"
	defaultName = "world"
)

var count = "0"

func main() {
	// Set up a connection to the server.
	conn, erro := grpc.Dial(address, grpc.WithInsecure())
	if erro != nil {
		log.Fatalf("did not connect: %v", erro)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	// name := defaultName
	// if len(os.Args) > 1 {
	// 	name = os.Args[1]
	// }
	stream, erro := c.SayHello(context.Background(), &pb.HelloRequest{Name: "2"})
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
		log.Println("Car Count: ", helloReply.Message)
	}
	
}

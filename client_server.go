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
	"fmt"
	"time"
	"net/http"
	"log"
	"io"
	"strings"
	"strconv"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"

	pbs "api.gateway/semanticContract"
	pbv "api.gateway/volumeContract"
	pbd "api.gateway/densityContract"
	pbg "api.gateway/gatewayContract"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc/reflection"
)

const (
	addressVolume = "volume-service:50051"
	addressDensity = "density-service:50050"
	addressSemantic = "semantic-service:50049"
)

var timestamp = ""

var density_state = "Lancar"

var semantic_state = ""

// server is used to implement helloworld.GreeterServer.
type server struct{}

func asClientDensity(id int32, resp chan string, ctx context.Context) {
	var withBlock = grpc.WithBlock()
	conn, erro := grpc.Dial(addressDensity, grpc.WithInsecure(), withBlock)
	if erro != nil {
		log.Printf("did not connect: %v", erro)
	}
	defer conn.Close()
	c := pbd.NewGreeterClient(conn)

	// Contact the server and print out its response.

	stream, erro := c.SayHello(context.Background(), &pbd.HelloRequest{Id: id})
	if erro != nil {
		log.Printf("could not greet: %v", erro)
	}
	for {
		select{
			case <- ctx.Done():
				fmt.Println("close client")
				return
			default:
				helloReply, erro := stream.Recv()
				if erro == io.EOF {
					return
				}
				if erro != nil {
					log.Printf("%v.ListFeatures(_) = _, %v", c, erro)
					return
				}
				resp <- helloReply.Response
				// log.Println("Density: ", helloReply.Response)
				time.Sleep(10 * time.Millisecond)
		}
	}
}

func asClientVolume(id int32, resp chan string, ctx context.Context) {
	////////// Client ////////////
	var withBlock = grpc.WithBlock()
	conn, erro := grpc.Dial(addressVolume, grpc.WithInsecure(), withBlock)
	if erro != nil {
		log.Printf("did not connect: %v", erro)
	}
	defer conn.Close()
	c := pbv.NewGreeterClient(conn)

	// Contact the server and print out its response.
	stream, erro := c.SayHello(context.Background(), &pbv.HelloRequest{Id: id})
	if erro != nil {
		log.Printf("could not greet: %v", erro)
	}

	for {
		select{
			case <- ctx.Done():
				fmt.Println("close client")
				return
			default:
				helloReply, erro := stream.Recv()
				if erro == io.EOF {
					return
				}
				if erro != nil {
					log.Printf("%v.ListFeatures(_) = _, %v", c, erro)
					return
				}
				resp <- strconv.FormatInt(int64(helloReply.Volume),10)
				// log.Println("Car Count: ", strconv.FormatInt(int64(helloReply.Volume),10))
				
				time.Sleep(10 * time.Millisecond)
		}
	}
}

func asClientSemantic(id int32, resp chan string, ctx context.Context) {
	var withBlock = grpc.WithBlock()
	conn, erro := grpc.Dial(addressSemantic, grpc.WithInsecure(), withBlock)
	if erro != nil {
		log.Printf("did not connect: %v", erro)
	}
	defer conn.Close()
	c := pbs.NewGreeterClient(conn)

	// Contact the server and print out its response.

	stream, erro := c.SayHello(context.Background(), &pbs.HelloRequest{Id: id})
	if erro != nil {
		log.Printf("could not greet: %v", erro)
	}
	for {
		select{
			case <- ctx.Done():
				fmt.Println("close client")
				return
			default:
				helloReply, erro := stream.Recv()
				if erro == io.EOF {
					return
				}
				if erro != nil {
					log.Printf("%v.ListFeatures(_) = _, %v", c, erro)
					return
				}
				resp <- helloReply.Response
				// log.Println("Semantic: ", helloReply.Response)
				time.Sleep(10 * time.Millisecond)
		}
	}
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(in *pbg.HelloRequest, stream pbg.Greeter_SayHelloServer) error {
	stream.SendHeader(metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-stream"))
	ctx := stream.Context()
	response := make(chan string)
	if strings.ToLower(in.Type) == "volume" {
		go asClientVolume(in.Id, response, ctx)
	} else if strings.ToLower(in.Type) == "density" {
		go asClientDensity(in.Id, response, ctx)
	} else if strings.ToLower(in.Type) == "semantic" {
		go asClientSemantic(in.Id, response, ctx)
	}
	for {
		select{
			case <- ctx.Done():
				fmt.Println("close context")
				return nil
			default:
				resp := <- response
				helloReply := &pbg.HelloReply{Response: resp}
				if err := stream.Send(helloReply); err != nil {
					return err
				}
				// fmt.Println("response : ",resp)
		}
	}
	stream.SetTrailer(metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-stream"))
	return nil
}

func asServer() {
	//////////// Server //////////////
	port := 8080
	s := grpc.NewServer()
	pbg.RegisterGreeterServer(s, &server{})
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
		grpclog.Printf("failed starting http server: %v", err)
	}
}

func main() {
	asServer()
}
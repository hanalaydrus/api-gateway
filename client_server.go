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
	// "io"
	"strings"
	"strconv"
	"encoding/json"
	"sync"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	// "google.golang.org/grpc/reflection"

	pbs "api.gateway/semanticContract"
	pbv "api.gateway/volumeContract"
	pbd "api.gateway/densityContract"
	pbg "api.gateway/gatewayContract"
)

const (
	addressVolume = "volume-service:50051"
	addressDensity = "density-service:50050"
	addressSemantic = "semantic-service:50049"
)

var connection_count = 0

// server is used to implement helloworld.GreeterServer.
type server struct{}

func asClientDensity(id int32, resp chan string, ctx context.Context, wg sync.WaitGroup) {
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
				fmt.Println("close density client")
				stream.CloseSend()
				conn.Close()
				close(resp)
				wg.Done()
				return
			default:
				helloReply, erro := stream.Recv()
				if erro != nil {
					log.Printf("%v.ListFeatures(_) = _, %v", c, erro)
					stream.CloseSend()
					conn.Close()
					close(resp)
					wg.Done()
					return
				}
				resp <- helloReply.Response
				// log.Println("Density: ", helloReply.Response)
				time.Sleep(10 * time.Millisecond)
		}
	}
}

func asClientVolume(id int32, resp chan string, ctx context.Context, wg sync.WaitGroup) {
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
				fmt.Println("close volume client")
				stream.CloseSend()
				conn.Close()
				close(resp)
				wg.Done()
				return
			default:
				helloReply, erro := stream.Recv()
				if erro != nil {
					fmt.Println("%v.ListFeatures(_) = _, %v", c, erro)
					stream.CloseSend()
					conn.Close()
					close(resp)
					wg.Done()
					return
				}
				resp <- strconv.FormatInt(int64(helloReply.Volume),10)
				// log.Println("Car Count: ", strconv.FormatInt(int64(helloReply.Volume),10))
				
				time.Sleep(10 * time.Millisecond)
		}
	}
}

func asClientSemantic(id int32, resp chan string, ctx context.Context, wg sync.WaitGroup) {
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
				fmt.Println("close semantic client")
				stream.CloseSend()
				conn.Close()
				close(resp)
				wg.Done()
				return
			default:
				helloReply, erro := stream.Recv()
				if erro != nil {
					fmt.Println("%v.ListFeatures(_) = _, %v", c, erro)
					stream.CloseSend()
					conn.Close()
					close(resp)
					wg.Done()
					return
				}
				resp <- helloReply.Response
				// log.Println("Semantic: ", helloReply.Response)
				time.Sleep(10 * time.Millisecond)
		}
	}
}

func serveRequest (stream pbg.Greeter_SayHelloServer, response chan string, ctx context.Context, wg sync.WaitGroup) {
	for {
		select {
			case <- ctx.Done():
				fmt.Println("close context")
				wg.Done()
				return
			default:
				resp := <- response
				helloReply := &pbg.HelloReply{Response: resp}
				if err := stream.Send(helloReply); err != nil {
					// fmt.Printf("Error: %s", err)
					continue
				}
		}
	}
}

func serveRequestAll (stream pbg.Greeter_SayHelloServer, responseVolume chan string, responseDensity chan string, responseSemantic chan string, ctx context.Context, wg sync.WaitGroup) {
	for {
		select {
			case <- ctx.Done():
				fmt.Println("close context")
				wg.Done()
				return
			default:
				type Message struct {
					Volume string `json:"volume"`
					Density string `json:"density"`
					Semantic string `json:"semantic"`
				}
				respVolume := <- responseVolume
				respDensity := <- responseDensity
				respSemantic := <- responseSemantic

				m := Message{Volume: respVolume, Density: respDensity, Semantic: respSemantic}
				b, err := json.Marshal(m)
				if err != nil {
			        fmt.Println("Error: %s", err)
			        continue
			    }

				helloReply := &pbg.HelloReply{Response: string(b)}
				if err := stream.Send(helloReply); err != nil {
					fmt.Println("Error: %s", err)
					continue
				}
		}
	}
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(in *pbg.HelloRequest, stream pbg.Greeter_SayHelloServer) error {
	var wg sync.WaitGroup
	stream.SendHeader(metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-stream"))
	stream.SetTrailer(metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-stream"))
	ctx := stream.Context()

	if strings.ToLower(in.Type) == "volume" {

		response := make(chan string)
		wg.Add(2)
		go asClientVolume(in.Id, response, ctx, wg)
		go serveRequest(stream, response, ctx, wg)
		wg.Wait()

	} else if strings.ToLower(in.Type) == "density" {

		response := make(chan string)
		wg.Add(2)
		go asClientDensity(in.Id, response, ctx, wg)
		go serveRequest(stream, response, ctx, wg)
		wg.Wait()

	} else if strings.ToLower(in.Type) == "semantic" {

		response := make(chan string)
		wg.Add(2)
		go asClientSemantic(in.Id, response, ctx, wg)
		go serveRequest(stream, response, ctx, wg)
		wg.Wait()

	} else if strings.ToLower(in.Type) == "all" {

		responseVolume := make(chan string)
		responseDensity := make(chan string)
		responseSemantic := make(chan string)

		wg.Add(4)
		go asClientVolume(in.Id, responseVolume, ctx, wg)
		go asClientDensity(in.Id, responseDensity, ctx, wg)
		go asClientSemantic(in.Id, responseSemantic, ctx, wg)
		go serveRequestAll(stream, responseVolume, responseDensity, responseSemantic, ctx, wg)
		wg.Wait()
	}

	return nil
}

func asServer() {
	//////////// Server //////////////
	port := 8080
	s := grpc.NewServer()
	pbg.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	// reflection.Register(s)

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
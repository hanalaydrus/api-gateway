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
	"flag"
	"fmt"
	"log"
	"net/http"
	// "net"
	"os"

	"google.golang.org/grpc"
	// "google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	// pb "google.golang.org/grpc/examples/helloworld/helloworld"
	// "google.golang.org/grpc/reflection"

	// "strings"

	library "../helloworld"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"golang.org/x/net/context"
)

var (
	enableTls       = flag.Bool("enable_tls", false, "Use TLS - required for HTTP2.")
	tlsCertFilePath = flag.String("tls_cert_file", "../misc/localhost.crt", "Path to the CRT/PEM file.")
	tlsKeyFilePath  = flag.String("tls_key_file", "../misc/localhost.key", "Path to the private key file.")
)

// const (
// 	port = ":50051"
// )

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *library.HelloRequest) (*library.HelloReply, error) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))
	grpc.SetTrailer(ctx, metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-unary"))

	return &library.HelloReply{Message: "Hello " + in.Name}, nil
}

// func main() {
// 	lis, err := net.Listen("tcp", port)
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}
// 	s := grpc.NewServer()
// 	pb.RegisterGreeterServer(s, &server{})
// 	// Register reflection service on gRPC server.
// 	reflection.Register(s)
// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }

func main() {
	flag.Parse()

	port := 9090
	if *enableTls {
		port = 9091
	}

	grpcServer := grpc.NewServer()
	library.RegisterGreeterServer(grpcServer, &server{})
	grpclog.SetLogger(log.New(os.Stdout, "exampleserver: ", log.LstdFlags))
	grpcweb.WithCorsForRegisteredEndpointsOnly(false)
	wrappedServer := grpcweb.WrapServer(grpcServer)
	handler := func(resp http.ResponseWriter, req *http.Request) {
		wrappedServer.ServeHTTP(resp, req)
	}

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(handler),
	}

	grpclog.Printf("Starting server. http port: %d, with TLS: %v", port, *enableTls)

	if *enableTls {
		if err := httpServer.ListenAndServeTLS(*tlsCertFilePath, *tlsKeyFilePath); err != nil {
			grpclog.Fatalf("failed starting http2 server: %v", err)
		}
	} else {
		if err := httpServer.ListenAndServe(); err != nil {
			grpclog.Fatalf("failed starting http server: %v", err)
		}
	}
}
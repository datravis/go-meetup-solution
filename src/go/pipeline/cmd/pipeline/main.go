package main

import (
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/datravis/go-meetup-solution/src/go/pipeline/pkg/srv"
	"github.com/datravis/go-meetup-solution/src/go/protogen"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	log.Println("Starting GRPC server")
	grpcServer := grpc.NewServer(opts...)
	protogen.RegisterPipelineServiceServer(grpcServer, srv.NewServer())
	grpcServer.Serve(lis)
}

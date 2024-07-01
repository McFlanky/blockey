package main

import (
	"fmt"
	"log"
	"net"

	"github.com/McFlanky/blocker/node"
	"github.com/McFlanky/blocker/proto"
	"google.golang.org/grpc"
)

func main() {
	node := node.NewNode()

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}
	proto.RegisterNodeServer(grpcServer, node)
	fmt.Println("node running on port", ":3000")
	grpcServer.Serve(ln)
}

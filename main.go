package main

import (
	"VideoStorageService/api"
	"VideoStorageService/gapi"
	"VideoStorageService/pb"
	"VideoStorageService/util"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load config file", err)
	}

	// server := api.NewServer()

	// err = server.Start(config.HTTPServerAddress)
	// if err != nil {
	// 	log.Fatal("cannot start server:", err)
	// }

	runGrpcServer(config)
}

func runGrpcServer(config util.Config) {

	server, err := gapi.NewServer()
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterVideoStorageServiceServer(grpcServer, server)
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener:", err)
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server:", err)
	}
}

func runGinServer(config util.Config) {
	server, err := api.NewServer()
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

package gapi

import (
	"VideoStorageService/pb"
	"VideoStorageService/util"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

type Server struct {
	pb.UnimplementedVideoStorageServiceServer
	router *gin.Engine
	s3     *s3.S3
	config util.Config
}

func NewServer() (*Server, error) {
	server := &Server{}
	return server, nil
}

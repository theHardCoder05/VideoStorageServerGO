package api

import (
	"VideoStorageService/util"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	s3     *s3.S3
	config util.Config
}

func NewServer() (*Server, error) {

	server := &Server{}
	router := gin.Default()
	v1 := router.Group("/v1")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
		Credentials: credentials.NewStaticCredentials(
			"AKIAYK7L25QDGS2HHJNJ",
			"k3++WiQ4JwzTSTQQDcjj3DWThJnnR5gmSL3mHXfS",
			""),
	})

	if err != nil {
		panic(err)
	}

	s3 := s3.New(sess)

	v1.GET("/health", server.HealthCheck)
	v1.POST("/files", server.Upload)
	v1.GET("/files/:id", server.GetContentById)
	v1.DELETE("/files/:id", server.RemoveContentById)
	v1.GET("/", server.Index)
	server.router = router
	server.s3 = s3
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load config file", err)
	}
	server.config = config
	return server, nil

}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

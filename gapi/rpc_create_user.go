package gapi

import (
	"VideoStorageService/pb"
	"context"
)

func (server *Server) CreateContent(ctx context.Context, req *pb.CreateContentRequest) (*pb.CreateContentResponse, error) {

	respose := &pb.CreateContentResponse{
		Content: &pb.Content{
			Fileid: req.Fileid,
			Name:   req.Name,
			Size:   req.Size,
		},
	}
	return respose, nil
}

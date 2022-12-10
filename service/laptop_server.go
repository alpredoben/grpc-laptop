package service

import (
	"context"
	pBuff "gocpu/pb"
	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LaptopServer struct {
	Store LaptopStore
}

func NewLaptopServer() *LaptopServer {
	return &LaptopServer{}
}

func (c *LaptopServer) CreateLaptop(ctx context.Context, req *pBuff.CreateLaptopRequest) (*pBuff.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("Receipt a create laptop requiest with id: %s", laptop.Id)

	if len(laptop.Id) > 0 {
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Laptop ID is not valid uuid : %v", err.Error())
		} 
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Cannot generate a new laptop ID: %v", err.Error())
		}
		laptop.Id = id.String()
	}

	//save laptop to store
}

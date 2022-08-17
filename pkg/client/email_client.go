package client

import (
	"context"
	"log"

	"github.com/Shulammite-Aso/filebox-service/pkg/proto"
	"google.golang.org/grpc"
)

type EmailServiceClient struct {
	Client proto.EmailServiceClient
}

func InitEmailServiceClient(url string) EmailServiceClient {
	cc, err := grpc.Dial(url, grpc.WithInsecure())

	if err != nil {
		log.Println("Could not connect:", err)
	}

	c := EmailServiceClient{
		Client: proto.NewEmailServiceClient(cc),
	}

	return c
}

func (c *EmailServiceClient) SendEmail(user, fileSent, fileSender string) (*proto.SendEmailResponse, error) {
	req := &proto.SendEmailRequest{
		User:       user,
		FileSent:   fileSent,
		FileSender: fileSender,
	}

	return c.Client.SendEmail(context.Background(), req)
}

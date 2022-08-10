package main

import (
	"log"
	"net"

	"github.com/Shulammite-Aso/filebox-service/pkg/config"
	"github.com/Shulammite-Aso/filebox-service/pkg/db"
	"github.com/Shulammite-Aso/filebox-service/pkg/proto"
	"github.com/Shulammite-Aso/filebox-service/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	log.Println("Product service listening on", c.Port)

	s := services.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	proto.RegisterFileboxServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}

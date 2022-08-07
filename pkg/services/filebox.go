package services

import (
	"context"
	"io/ioutil"
	"log"
	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Shulammite-Aso/filebox-service/pkg/db"
	"github.com/Shulammite-Aso/filebox-service/pkg/models"
	"github.com/Shulammite-Aso/filebox-service/pkg/proto"
	"github.com/Shulammite-Aso/filebox-service/pkg/utils"
)

type Server struct {
	H db.Handler
}

func (s *Server) SendFile(ctx context.Context, req *proto.SendFileRequest) (*proto.SuccessMessage, error) {

	// 1. Create a file using the fileName from request
	// 2. Write []byte from request into the file
	// 3. Upload the file to cloudinary
	// 4. Get public URL of downloaded file and set it as the value of file from the model
	// 5. Create the file entry in the DB

	err := ioutil.WriteFile(req.FileName, req.File, 0)

	if err != nil {
		return &proto.SuccessMessage{}, status.Errorf(codes.Internal, err.Error())
	}

	defer os.Remove(req.FileName)

	url, publicId, err := utils.UploadFileHelper(req.FileName)

	if err != nil {
		return &proto.SuccessMessage{}, status.Errorf(codes.Internal, err.Error())
	}

	fileEntry := models.FileEntry{
		Username:      req.Username,
		FileName:      req.FileName,
		File:          url,
		CloudPublicID: publicId,
	}

	res, err := s.H.Collection.InsertOne(context.Background(), fileEntry)
	if err != nil {
		return &proto.SuccessMessage{}, status.Errorf(codes.Internal, err.Error())
	}

	log.Println("create file entry:", res.InsertedID)

	return &proto.SuccessMessage{
		Message: "File uploaded",
	}, nil
}

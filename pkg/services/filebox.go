package services

import (
	"context"
	"io/ioutil"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Shulammite-Aso/filebox-service/pkg/client"
	"github.com/Shulammite-Aso/filebox-service/pkg/db"
	"github.com/Shulammite-Aso/filebox-service/pkg/models"
	"github.com/Shulammite-Aso/filebox-service/pkg/proto"
	"github.com/Shulammite-Aso/filebox-service/pkg/utils"
)

type Server struct {
	proto.UnimplementedFileboxServiceServer
	H        db.Handler
	EmailSvc client.EmailServiceClient
}

func (s *Server) SendFile(ctx context.Context, req *proto.SendFileRequest) (*proto.SuccessMessage, error) {

	// 1. Create a file using the fileName from request
	// 2. Write []byte from request into the file
	// 3. Upload the file to cloudinary
	// 4. Get public URL and PublicID of downloaded file and set it as the values of file & CloudPublicID from the model
	// 5. Create the file entry in the DB

	var fileEntry models.FileEntry

	result := s.H.Collection.FindOne(context.Background(), bson.D{{"username", req.Username}, {"fileName", req.FileName}})

	if err := result.Decode(&fileEntry); err != mongo.ErrNoDocuments {
		return &proto.SuccessMessage{}, status.Errorf(codes.InvalidArgument, "You already own a file with this name, please rename file")
	}

	err := ioutil.WriteFile(req.FileName, req.File, 0)

	if err != nil {
		return &proto.SuccessMessage{}, status.Errorf(codes.Internal, err.Error())
	}

	defer os.Remove(req.FileName)

	url, publicId, err := utils.UploadFileHelper(req.FileName)

	if err != nil {
		return &proto.SuccessMessage{}, status.Errorf(codes.Internal, err.Error())
	}

	fileEntry = models.FileEntry{
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

func (s *Server) UpdateFile(ctx context.Context, req *proto.UpdateFileRequest) (*proto.SuccessMessage, error) {

	// 1. Do step one to four from SendFile function
	// 2. Update the matching file entry in the DB
	// 3. Delete the previous file from cloudinary

	err := ioutil.WriteFile(req.FileName, req.File, 0)

	if err != nil {
		return &proto.SuccessMessage{}, status.Errorf(codes.Internal, err.Error())
	}

	defer os.Remove(req.FileName)

	url, publicId, err := utils.UploadFileHelper(req.FileName)

	if err != nil {
		return &proto.SuccessMessage{}, status.Errorf(codes.Internal, err.Error())
	}

	var fileEntry models.FileEntry

	filter := bson.D{{"username", req.Username}, {"fileName", req.FileName}}
	update := bson.D{{"$set", bson.D{{"file", url}, {"CloudPublicID", publicId}}}}

	result := s.H.Collection.FindOneAndUpdate(context.Background(), filter, update)

	if err := result.Decode(&fileEntry); err != nil {
		if err == mongo.ErrNoDocuments {
			return &proto.SuccessMessage{}, status.Errorf(codes.InvalidArgument, "No such file in your box")
		}
		return &proto.SuccessMessage{}, status.Errorf(codes.Internal, err.Error())
	}

	err3 := utils.RemoveFileHelper(fileEntry.CloudPublicID)

	if err3 != nil {
		log.Printf("could not delete file %v from cloudinary after update: %v \n", fileEntry.CloudPublicID, err3)
	}

	return &proto.SuccessMessage{
		Message: "File updated",
	}, nil
}

func (s *Server) GetFile(ctx context.Context, req *proto.GetFileRequest) (*proto.GetFileResponse, error) {

	var fileEntry models.FileEntry

	result := s.H.Collection.FindOne(context.Background(), bson.D{{"username", req.Username}, {"fileName", req.FileName}})

	if err := result.Decode(&fileEntry); err != nil {
		if err == mongo.ErrNoDocuments {
			return &proto.GetFileResponse{}, status.Errorf(codes.InvalidArgument, "No such file in your box")
		}
		return &proto.GetFileResponse{}, status.Errorf(codes.Internal, err.Error())
	}

	return &proto.GetFileResponse{FileURL: fileEntry.File}, nil
}

func (s *Server) GetListOfAllFiles(ctx context.Context, req *proto.GetListOfAllFilesRequest) (*proto.GetListOfAllFilesResponse, error) {

	var results []models.FileEntry
	var allFiles []string

	cur, err := s.H.Collection.Find(context.Background(), bson.D{{"username", req.Username}})

	if err != nil {
		log.Fatal(err)
	}

	if err := cur.All(context.Background(), &results); err != nil {
		return &proto.GetListOfAllFilesResponse{}, status.Errorf(codes.Internal, err.Error())
	}

	for _, result := range results {
		allFiles = append(allFiles, result.File)
	}

	return &proto.GetListOfAllFilesResponse{AllFiles: allFiles}, nil
}

func (s *Server) DeleteFile(ctx context.Context, req *proto.DeleteFileRequest) (*proto.SuccessMessage, error) {

	var fileEntry models.FileEntry

	result := s.H.Collection.FindOneAndDelete(context.Background(), bson.D{{"username", req.Username}, {"fileName", req.FileName}})

	if err := result.Decode(&fileEntry); err != nil {
		if err == mongo.ErrNoDocuments {
			return &proto.SuccessMessage{}, status.Errorf(codes.InvalidArgument, "No such file in your box")
		}
		return &proto.SuccessMessage{}, status.Errorf(codes.Internal, err.Error())
	}

	err := utils.RemoveFileHelper(fileEntry.CloudPublicID)

	if err != nil {
		log.Printf("could not delete file %v from cloudinary: %v \n", fileEntry.CloudPublicID, err)
	}

	return &proto.SuccessMessage{
		Message: "File deleted",
	}, nil

}

func (s *Server) SendFileToPerson(ctx context.Context, req *proto.SendFileToPersonRequest) (*proto.SuccessMessage, error) {

	//
	// 1. Create a file using the fileName from request
	// 2. Write []byte from request into the file
	// 3. Upload the file to cloudinary
	// 4. Get public URL and PublicID of downloaded file and set it as the values of file & CloudPublicID from the model
	// 5. Create the file entry in the DB

	var fileEntry models.FileEntry

	result := s.H.Collection.FindOne(context.Background(), bson.D{{"username", req.Username}, {"fileName", req.FileName}})

	if err := result.Decode(&fileEntry); err != mongo.ErrNoDocuments {
		return &proto.SuccessMessage{}, status.Errorf(codes.InvalidArgument, "this user already own a file with this name, please rename file")
	}

	err := ioutil.WriteFile(req.FileName, req.File, 0)

	if err != nil {
		return &proto.SuccessMessage{}, status.Errorf(codes.Internal, err.Error())
	}

	defer os.Remove(req.FileName)

	url, publicId, err := utils.UploadFileHelper(req.FileName)

	if err != nil {
		return &proto.SuccessMessage{}, status.Errorf(codes.Internal, err.Error())
	}

	fileEntry = models.FileEntry{
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

	// Call email service to send email to receiver
	if _, err := s.EmailSvc.SendEmail(req.Username, req.FileName, req.SenderUsername); err != nil {
		log.Println(err)
	}

	return &proto.SuccessMessage{
		Message: "File uploaded",
	}, nil
}

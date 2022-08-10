package utils

import (
	"context"
	"time"

	"github.com/Shulammite-Aso/filebox-service/pkg/config"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func UploadFileHelper(file string) (string, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c, err := config.LoadConfig()

	if err != nil {
		return "", "", err
	}

	cld, err := cloudinary.NewFromParams(c.CloudName, c.CloudApiKey, c.CloudApiSecret)
	if err != nil {
		return "", "", err
	}

	resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{Folder: c.CloudFolder})
	if err != nil {
		return "", "", err
	}

	return resp.SecureURL, resp.PublicID, nil
}

func RemoveFileHelper(filePublicId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c, err := config.LoadConfig()

	if err != nil {
		return err
	}

	cld, err := cloudinary.NewFromParams(c.CloudName, c.CloudApiKey, c.CloudApiSecret)
	if err != nil {
		return err
	}

	_, err2 := cld.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: filePublicId})
	if err2 != nil {
		return err2
	}

	return nil
}

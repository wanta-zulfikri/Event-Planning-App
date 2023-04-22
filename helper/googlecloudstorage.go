package helper

import (
	"context"
	"io"
	"log"
	"mime/multipart"
	"time"

	"github.com/wanta-zulfikri/Event-Planning-App/config/common"

	"cloud.google.com/go/storage"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

type StorageGCPConfig struct {
	GCPClient  *storage.Client
	ProjectID  string
	BucketName string
	Path       string
}

func UploadImage(c echo.Context, file *multipart.FileHeader) (string, error) {
	if file == nil {
		return "", nil
	}
	image, err := file.Open()
	if err != nil {
		return "", err
	}
	defer image.Close()
	sgcp := StorageGCPConfig{
		GCPClient:  InitGCPClient(),
		ProjectID:  common.ProjectID,
		BucketName: common.BucketName,
		Path:       common.Path,
	}

	imageURL, err := sgcp.UploadFile(image, file.Filename)
	if err != nil {
		return "", err
	}
	return imageURL, nil
}

func InitGCPClient() *storage.Client {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(common.Credential))
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (s *StorageGCPConfig) UploadFile(file io.Reader, fileName string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	wc := s.GCPClient.Bucket(s.BucketName).Object(s.Path + fileName).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return "", err
	}
	if err := wc.Close(); err != nil {
		return "", err
	}
	image := "https://storage.googleapis.com/" + s.BucketName + "/" + s.Path + fileName
	return image, nil
}

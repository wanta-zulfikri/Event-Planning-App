package helper

import (
	"context"
	"io"
	"mime/multipart"
	"time"

	"cloud.google.com/go/storage"
	// "github.com/cloudinary/cloudinary-go/v2" 
	// "github.com/cloudinary/cloudinary-go/v2/api/uploader"

)

type StorageGCP struct {
	ClG        *storage.Client
	ProjectID  string
	BucketName string
	Path       string
}

func (s *StorageGCP) UploadFile(file multipart.File, fileName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	wc := s.ClG.Bucket(s.BucketName).Object(s.Path + fileName).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return err
	}

	if err := wc.Close(); err != nil {
		return err
	}
	return nil
} 

// func UploadFile(fileContents interface{}, path string) ([]string, error) {
// 	var urls []string 
// 	switch cnv := fileContents.(type) {
// 	case []multipart.File: 
// 		for _, content := range cnv {
// 			uploadResult, err := getLink(content, path) 
// 			if err != nil {
// 			return nil, err
// 		} 
// 		urls = append(urls, uploadResult)
// 	}
// case multipart.File: 
// 	uploadResult, err := getLink(cnv, path)
// 	if err != nil {
// 		return nil, err
// 	}
// 	urls = append(urls, uploadResult)
// }
// return urls, nil
// }
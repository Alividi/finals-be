package service

import (
	"context"
	"finals-be/app/shared/dto"
	"finals-be/internal/config"
	"finals-be/internal/constants"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Service struct {
	cfg    *config.Config
	client *s3.Client
}

func NewS3Service(config *config.Config, client *s3.Client) *S3Service {
	return &S3Service{
		cfg:    config,
		client: client,
	}
}

func (s *S3Service) GetImageURL(fileName string, fileType string) string {
	return fmt.Sprintf("%s/%s/%s/%s", s.cfg.AWS.BaseURL, constants.S3_ROOT_PATH, fileType, fileName)
}

func (s *S3Service) UploadFile(ctx context.Context, fileName string, fileType string, userID int64, file io.Reader, contentType string) (string, error) {
	timestamp := time.Now().Unix()
	uniqueFileName := fmt.Sprintf("%s_%d_%s", fileType, timestamp, fileName)

	filePath := fmt.Sprintf("%s/%s/%s", constants.S3_ROOT_PATH, fileType, uniqueFileName)
	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      &s.cfg.AWS.BucketName,
		Key:         &filePath,
		Body:        file,
		ContentType: &contentType,
	})

	if err != nil {
		return "", err
	}

	return s.GetImageURL(uniqueFileName, fileType), nil
}

func (s *S3Service) UploadFiles(ctx context.Context, fileType string, userID int64, files []dto.FileUpload) ([]string, error) {
	var urls []string
	var mu sync.Mutex
	var wg sync.WaitGroup
	errChan := make(chan error, len(files))

	for _, f := range files {
		wg.Add(1)
		go func(file dto.FileUpload) {
			defer wg.Done()
			url, err := s.UploadFile(ctx, file.Name, fileType, userID, file.Reader, file.ContentType)
			if err != nil {
				errChan <- err
				return
			}
			mu.Lock()
			urls = append(urls, url)
			mu.Unlock()
		}(f)
	}

	wg.Wait()
	close(errChan)

	if len(errChan) > 0 {
		return nil, <-errChan
	}

	return urls, nil
}

func (s *S3Service) DeleteFile(ctx context.Context, fileName string, fileType string) error {
	filePath := fmt.Sprintf("%s/%s/%s", constants.S3_ROOT_PATH, fileType, fileName)

	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: &s.cfg.AWS.BucketName,
		Key:    &filePath,
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *S3Service) GetFileNameFromURL(url string) (string, error) {
	if url == "" {
		return "", fmt.Errorf("url is empty")
	}

	parts := strings.Split(url, "/")

	if len(parts) < 5 && s.cfg.AWS.BaseURL != fmt.Sprintf("%s/%s", parts[0], parts[1]) {
		return "", fmt.Errorf("invalid URL format")
	}

	fileName := parts[len(parts)-1]

	return fileName, nil
}

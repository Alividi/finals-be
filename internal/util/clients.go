package util

import (
	"finals-be/internal/connection"

	"firebase.google.com/go/v4/messaging"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Clients struct {
	DB       *connection.SQLServerConnectionManager
	Message  *messaging.Client
	S3Client *s3.Client
}

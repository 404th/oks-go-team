package image

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
)

func ImageUpload(file []byte) (*s3manager.UploadOutput, error) {
	s3Config := &aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWSS3ACCESSLOGIN"), os.Getenv("AWSS3SECRETACCESSKEY"), ""),
	}
	s3Session, err := session.NewSession(s3Config)
	if err != nil {
		return nil, err
	}

	uploader := s3manager.NewUploader(s3Session)

	input := &s3manager.UploadInput{
		Bucket:      aws.String(os.Getenv("AWSS3BUCKETNAME")),                   // bucket's name
		Key:         aws.String(fmt.Sprintf("images/%s.jpg", uuid.NewString())), // files destination location
		Body:        bytes.NewReader(file),                                      // content of the file
		ContentType: aws.String("image/jpeg"),                                   // content type
	}
	return uploader.UploadWithContext(context.Background(), input)
}

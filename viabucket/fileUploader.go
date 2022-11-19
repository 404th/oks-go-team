package main

import (
	"context"
	"log"
	"os"

	"github.com/minio/minio-go/pkg/credentials"
	"github.com/minio/minio-go/v7"
)

func uploadFile() {
	// AWSS3ACCESSLOGIN=AKIA6FYS75ENMPWHIZJX
	// AWSS3SECRETACCESSKEY=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
	// AWSS3BUCKETNAME=storageforfile
	// AWSS3BUCKETLOCATION=ap-southeast-1

	ctx := context.Background()
	endpoint := "play.min.io"
	accessKeyID := os.Getenv("AWSS3ACCESSLOGIN")
	secretAccessKey := os.Getenv("AWSS3SECRETACCESSKEY")
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Make a new bucket called mymusic.
	bucketName := "mymusic"
	location := "us-east-1"

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	// Upload the zip file
	objectName := "golden-oldies.zip"
	filePath := "/tmp/golden-oldies.zip"
	contentType := "application/zip"

	// Upload the zip file with FPutObject
	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
}

package fileUploader

import (
	"context"
	"log"
	"mime/multipart"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func UploadFile(fl *multipart.FileHeader) {
	// AWSS3ACCESSLOGIN=AKIA6FYS75ENMPWHIZJX
	// AWSS3SECRETACCESSKEY=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
	// AWSS3BUCKETNAME=storageforfile
	// AWSS3BUCKETLOCATION=ap-southeast-1

	ctx := context.Background()
	endpoint := "minioaccesstraining-974460545306.s3-accesspoint.ap-southeast-1.amazonaws.com"
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

	// Make a new bucket called 'storageforfile' if not exists
	bucketName := os.Getenv("AWSS3BUCKETNAME")
	location := os.Getenv("AWSS3BUCKETLOCATION")

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
	objectName := fl.Filename
	filePath := "/tmp"
	contentType := "application/jpg"

	// Upload the zip file with FPutObject
	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
}

package common

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadImageToS3(fileBuffer []byte, fileName string) (string, error) {
	accessKey := os.Getenv("SUPABASE_ACCESS_KEY")
	secretKey := os.Getenv("SUPABASE_SECRET_KEY")
	bucketName := os.Getenv("SUPABASE_BUCKET")
	endPoint := os.Getenv("SUPABASE_ENPOINT")
	region := "ap-southeast-1"

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			accessKey, // Đọc từ biến môi trường
			secretKey, // Đọc từ biến môi trường
			"",
		),
		Endpoint:         aws.String(endPoint),
		S3ForcePathStyle: aws.Bool(true),
	})

	if err != nil {
		return "", fmt.Errorf("failed to create AWS session: %w", err)
	}

	uploader := s3manager.NewUploader(sess)

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(fileBuffer),
	})

	if err != nil {
		return "", fmt.Errorf("failed to upload image: %w", err)
	}

	log.Printf("Upload result: %v", result)

	// fileURL := result.Location
	supabaseURL := os.Getenv("SUPABASE_URL")
	fileURL := fmt.Sprintf("%s/storage/v1/object/public/%s/%s", supabaseURL, bucketName, fileName)
	return fileURL, nil
}

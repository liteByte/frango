package frango

import (
	"io"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func PrintS3Buckets(region string) {
	svc := s3.New(session.New(&aws.Config{Region: aws.String(region)}))

	result, err := svc.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		log.Println("Failed to list buckets: ", err)
		return
	}

	log.Println("Buckets: ")

	for _, bucket := range result.Buckets {
		log.Printf("%s : %s\n", aws.StringValue(bucket.Name), bucket.CreationDate)
	}
}

func UploadImageToS3(file io.Reader, region, bucketName, fileName string) string {
	uploader := s3manager.NewUploader(session.New(&aws.Config{Region: aws.String(region)}))

	result, err := uploader.Upload(&s3manager.UploadInput{
		Body:   file,
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
	})

	if err != nil {
		log.Fatalln("Failed to upload file to S3: ", err)
	}

	return result.Location
}

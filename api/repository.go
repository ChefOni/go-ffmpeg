package api

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ChefOni/go-ffmpeg/database"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"gorm.io/gorm"
)


type S3Provider struct {
	BucketName string
}


func WriteLog(ctx context.Context, db *gorm.DB, logData map[string]interface{}) error{
	log := database.DBLogs{
		Date:                 logData["date"].(string),
		StorageBucketProvider: logData["provider"].(string),
		UserIP:               logData["userIP"].(string),
		RegisteredUser:       logData["isRegistered"].(bool),
		CreatedAt:            logData["createdAt"].(string),
		Status:               logData["status"].(bool),
		Size:                 logData["size"].(uint),
	}

	if err := db.Create(&log).Error; err != nil {
		return fmt.Errorf("error writing log to db: %v", err)
	}

	return nil
}


// Connect to storage bucket based on the provider
func ConnectStorageBucket(provider CloudProvider) error {
	return provider.Upload(context.Background(), "test_file_path")
}

func (s *S3Provider) Upload(ctx context.Context, filePath string) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		return err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	uploader := s3.New(sess)
	_, err = uploader.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(filePath),
		Body:   file,
	})
	if err != nil {
		return err
	}

	log.Println("File uploaded to S3 successfully")
	return nil
}


// UPLOAD TO CLOUDFLARE
func UploadtoCloudFlare(){}


//UPLOAD TO CLOUDINARY
func UploadToCloudinary(){}


// SINGLE UPLOAD FILE THAT RECEIVES A FILE PERFORM THE CHECK VIDEO SIZE COMPRESSES THE VIDEO , CHECKS FOR THE TYPE OF CLOUD PROVIDER USES THE UPLOAD FUNCTION AND WRITES TO THE DB 
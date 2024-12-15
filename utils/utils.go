package utils

import (
	"context"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	_ "github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func CloudinaryCredentials() (*cloudinary.Cloudinary, context.Context) {
    cld, _ := cloudinary.New()
    cld.Config.URL.Secure = true
    ctx := context.Background()
    return cld, ctx
}

func UploadImageToCloudinary(cld *cloudinary.Cloudinary, ctx context.Context) {
    resp, err := cld.Upload.Upload(ctx, "https://cloudinary-devs.github.io/cld-docs-assets/assets/images/butterfly.jpeg", uploader.UploadParams{
        PublicID:       "quickstart_butterfly",
        UniqueFilename: api.Bool(false),
        Overwrite:      api.Bool(true)})

    if err != nil {
        fmt.Println("error")
    }

    fmt.Println("****2. Upload an image****\nDelivery URL:", resp.SecureURL)
}

// CHECK CONNECTIONS OF THE DIFFERENT STORAGE BUCKET
func ConnectStorageBucket(){

}

// MAKE SURE THE HANDLER DOES NOT TKE MORE THAN THE REQUIRED VIDEO SIZE
func CheckVideoSize(){}

//USE FFMPEG TO COMPRESS THE VIDEO
func CompressVideo(){}

// UPLOAD TO AWS S3 BUCKET
func UploadtoS3(){

}

// UPLOAD TO CLOUDFLARE
func UploadtoCloudFlare(){}


//UPLOAD TO CLOUDINARY
func UploadToCloudinary(){}




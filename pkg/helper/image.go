package helper

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/spf13/viper"
)

func ImageUploadHelper(input interface{}) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cld, err := cloudinary.NewFromParams(viper.GetString("cloudinary.cloudName"), viper.GetString("cloudinary.apiKey"), viper.GetString("cloudinary.apiSecret"))
	if err != nil {
		return "", err
	}

	uploadParam, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{Folder: viper.GetString("cloudinary.uploadFolder")})
	if err != nil {
		return "", err
	}
	return uploadParam.SecureURL, nil
}

func ValidateImage(file multipart.File) error {
	buff := make([]byte, 512)
	if _, err := file.Read(buff); err != nil {
		return err
	}

	filetype := http.DetectContentType(buff)

	switch filetype {
	case "image/jpeg", "image/jpg":
		return nil
	case "image/gif":
		return nil
	case "image/png":
		return nil
	case "application/pdf":
		return nil
	default:
		return fmt.Errorf("uknown file")
	}
}

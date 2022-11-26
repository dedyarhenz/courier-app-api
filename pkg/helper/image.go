package helper

import (
	"context"
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

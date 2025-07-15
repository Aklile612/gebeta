package media

import (
	"context"
	"mime/multipart"

	"github.com/aklile/recipe-backend/internal/config"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)


func UploadImage(file multipart.File, fileHeader *multipart.FileHeader)(string,error){
	cloudName, apiKey, apiSecret := config.CLOUDINARYCREDINTIALS()
	cld,err:= cloudinary.NewFromParams(cloudName,apiKey,apiSecret)

	if err!= nil{
		return "",err
	}

	uploadResult,err:= cld.Upload.Upload(context.Background(),file,uploader.UploadParams{Folder: "recipes"})

	if err != nil{
		return "",err
	}

	return uploadResult.SecureURL,nil
}
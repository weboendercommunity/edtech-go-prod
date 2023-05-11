package fileupload

import (
	"context"
	"errors"
	"mime/multipart"
	"os"

	"edtech.id/pkg/utils"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

type FileUpload interface {
	Upload(file multipart.FileHeader) (*string, error)
	Delete(file string) (*string, error)
}

type FileUploadImpl struct {
}

// Delete implements Image
func (fileImpl *FileUploadImpl) Delete(file string) (*string, error) {
	cloud, err := cloudinary.NewFromURL("cloudinary://" + os.Getenv("CLOUDINARY_API_KEY") + ":" + os.Getenv("CLOUDINARY_SECRET_KEY") + "@" + os.Getenv("CLOUDINARY_CLOUD_NAME"))

	if err != nil {
		return nil, err
	}

	var ctx = context.Background()

	fileName := utils.GetFileName(file)

	resourceType := "image"

	if utils.IsVideo(file) {
		resourceType = "video"
	}

	response, err := cloud.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID:     fileName,
		ResourceType: resourceType,
	})

	if err != nil {
		return nil, err
	}

	return &response.Result, nil
}

// Upload implements Image
func (*FileUploadImpl) Upload(image multipart.FileHeader) (*string, error) {
	cloud, err := cloudinary.NewFromURL("cloudinary://" + os.Getenv("CLOUDINARY_API_KEY") + ":" + os.Getenv("CLOUDINARY_SECRET_KEY") + "@" + os.Getenv("CLOUDINARY_CLOUD_NAME"))

	if err != nil {
		return nil, err
	}

	var ctx = context.Background()

	binary, err := image.Open()

	if err != nil {
		return nil, err
	}

	if binary == nil {
		return nil, errors.New("image is empty")
	}

	defer binary.Close()
	uploadResult, err := cloud.Upload.Upload(
		ctx,
		binary,
		uploader.UploadParams{
			PublicID: uuid.New().String(),
		},
	)

	if err != nil {
		return nil, err
	}

	return &uploadResult.SecureURL, nil
}

func NewFileUpload() FileUpload {
	return &FileUploadImpl{}
}

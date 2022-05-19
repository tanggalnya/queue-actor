package gdrive

import (
	"context"

	"google.golang.org/api/drive/v3"
)

type Client interface {
	CreateFolder(ctx context.Context, name string, path string) error
}

type Config struct {
	FileSecretLocation string
}

type googleDriveClient struct {
	service *drive.Service
}

func (g googleDriveClient) CreateFolder(ctx context.Context, name string, path string) error {
	//TODO implement me
	panic("implement me")
}

func New(gd drive.Service) *googleDriveClient {
	return &googleDriveClient{service: &gd}
}

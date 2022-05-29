package gdrive

import (
	"context"
	"log"

	"google.golang.org/api/drive/v3"
)

const (
	ownerEmail           = "tanggalnyacom@gmail.com"
	permissionType       = "user"
	permissionRole       = "writer"
	googleMimeTypeFolder = "application/vnd.google-apps.folder"
)

type Client interface {
	CreateFolder(ctx context.Context, name string, parentID string) (ID string, err error)
	CreateFile(ctx context.Context, name string, folderID string, mimeType string) (*drive.File, error)
}

type Config struct {
	FileSecretLocation string
}

type googleDriveClient struct {
	service *drive.Service
}

func (g googleDriveClient) CreateFolder(ctx context.Context, name string, parentID string) (string, error) {
	fileMetadata := &drive.File{
		MimeType: googleMimeTypeFolder,
		Name:     name,
		Parents:  []string{parentID},
	}
	folder, err := g.service.Files.Create(fileMetadata).Fields("id").Do()
	if err != nil {
		return "", err
	}
	log.Printf("Folder ID: %s", folder.Id)

	permission := &drive.Permission{
		EmailAddress: ownerEmail,
		Role:         permissionRole,
		Type:         permissionType,
		PendingOwner: true,
	}
	permissionCall, err := g.service.Permissions.Create(folder.Id, permission).Do()
	if err != nil {
		return "", err
	}
	log.Printf("Permision ID: %s", permissionCall.Id)

	return folder.Id, nil
}

func (g googleDriveClient) CreateFile(ctx context.Context, name string, folderID string, mimeType string) (*drive.File, error) {
	fileMetadata := &drive.File{Name: name, Parents: []string{folderID}, MimeType: mimeType}
	file, err := g.service.Files.Create(fileMetadata).Fields("id", "name").Do()
	if err != nil {
		return nil, err
	}
	return file, nil
}

func New(gd drive.Service) *googleDriveClient {
	return &googleDriveClient{service: &gd}
}

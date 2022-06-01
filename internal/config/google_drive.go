package config

import "github.com/tanggalnya/queue-actor/pkg/gdrive"

var googleDrive gdrive.Config

func GoogleDriveConfig() gdrive.Config {
	return googleDrive
}

func initGoogleDriveConfig() {
	googleDrive = gdrive.Config{
		FileSecretLocation: getStringOrPanic("GOOGLE_DRIVE_FILE_SECRET_LOCATION"),
	}
}

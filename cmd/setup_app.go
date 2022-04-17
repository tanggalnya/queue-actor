package cmd

import (
	"github.com/tanggalnya/queue-actor/internal/config"
)

func InitAppBase(envPath string) error {
	if err := config.LoadApp(envPath); err != nil {
		return err
	}
	return nil
}

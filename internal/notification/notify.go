package notification

import (
	"github.com/gen2brain/beeep"
)

// NotifyNewStatus sends a desktop notification about a new tracking status
func NotifyNewStatus() {
	err := beeep.Notify("Yeni Durum", "Kargo ilerledi", "goph.png")
	if err != nil {
		panic(err)
	}
}

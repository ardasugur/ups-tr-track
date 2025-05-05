package scheduler

import (
	"upsTrack/internal/tracking"
	"upsTrack/internal/ui"
)

// CreateTrackingJob defines the periodic tracking check operation
func CreateTrackingJob(ts *tracking.TrackingService) func() {
	return func() {
		ui.ClearConsole()
		ts.Logger.Printf("Takip numarası kontrol ediliyor: %s", ts.TrackingID)
		trackingData := ts.TrackPackage()
		ui.PrintTrackingData(trackingData, ts.Styles)
	}
}

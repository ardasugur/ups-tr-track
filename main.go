package main

import (
	"github.com/go-co-op/gocron"
	"io"
	"log"
	"os"
	"time"

	"upsTrack/internal/scheduler"
	"upsTrack/internal/tracking"
	"upsTrack/internal/ui"
)

// createLogger initializes and returns a logger for the application
func createLogger() *log.Logger {
	logger := log.New(os.Stdout, "[UPS Takip] ", log.LstdFlags)
	logger.SetOutput(io.Discard) // Hide logs by redirecting to io.Discard
	return logger
}

func main() {
	// Create logger and styles
	logger := createLogger()
	styles := ui.NewTrackingStyles()

	// Get tracking ID from user
	trackingID := ui.PrintPrompt()

	// Create tracking service
	trackingService := tracking.NewTrackingService(logger, styles, trackingID)

	// Create a scheduler
	s := gocron.NewScheduler(time.UTC)

	// Define and set up the tracking job
	job := scheduler.CreateTrackingJob(trackingService)

	// Run the job immediately once for initial tracking
	job()

	// Schedule the job to run every minute
	s.Every(1).Minute().Do(job)

	// Start the scheduler in a non-blocking manner
	s.StartAsync()

	logger.Println("Takip zamanlayıcısı başlatıldı. Çıkmak için Ctrl+C tuşlarına basın.")

	// Keep the program running
	select {}
}

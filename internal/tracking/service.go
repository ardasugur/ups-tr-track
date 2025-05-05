package tracking

import (
	"fmt"
	"log"
	"strings"
	"upsTrack/internal/models"

	"github.com/gocolly/colly"
	"upsTrack/internal/notification"
)

// TrackingService handles UPS package tracking operations
type TrackingService struct {
	Logger             *log.Logger
	TrackingLink       string
	TrackingID         string
	PreviousTrackCount int
	Styles             models.TrackingStyles
	Collector          *colly.Collector
}

// NewTrackingService creates a new tracking service instance
func NewTrackingService(logger *log.Logger, styles models.TrackingStyles, trackingID string) *TrackingService {
	trackingLink := fmt.Sprintf("https://ups.com.tr/WaybillSorgu.aspx?Waybill=%s", trackingID)

	return &TrackingService{
		Logger:             logger,
		PreviousTrackCount: 0,
		TrackingLink:       trackingLink,
		TrackingID:         trackingID,
		Styles:             styles,
		Collector:          initCollector(logger),
	}
}

// initCollector creates and configures a colly collector
func initCollector(logger *log.Logger) *colly.Collector {
	collector := colly.NewCollector(
		colly.AllowedDomains("ups.com.tr"),
		colly.AllowURLRevisit(),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
	)

	// Set up debug logging for this visit
	collector.OnResponse(func(r *colly.Response) {
		logger.Printf("Visited: %s", r.Request.URL)
		logger.Printf("Response size: %d bytes", len(r.Body))
	})

	// Handle errors
	collector.OnError(func(r *colly.Response, err error) {
		logger.Printf("Request URL: %s failed with error: %v", r.Request.URL, err)
	})

	return collector
}

// TrackPackage retrieves tracking information for a UPS package
func (ts *TrackingService) TrackPackage() []models.Tracking {
	var trackingData []models.Tracking
	processedDates := make(map[string]bool)

	// Setup HTML callbacks for this request
	ts.setupHTMLCallbacks(processedDates, &trackingData)

	// Make the HTTP request
	err := ts.Collector.Visit(ts.TrackingLink)
	if err != nil {
		ts.Logger.Printf("Error visiting the tracking link: %v", err)
	}

	// Handle tracking results
	ts.processTrackingResults(trackingData)

	return trackingData
}

// setupHTMLCallbacks configures the HTML parsing callbacks for the collector
func (ts *TrackingService) setupHTMLCallbacks(processedDates map[string]bool, trackingData *[]models.Tracking) {
	ts.Collector.OnHTML("span[id*='DataListSonIslem_ctl']", func(e *colly.HTMLElement) {
		id := e.Attr("id")
		text := strings.TrimSpace(e.Text)

		// Skip header row spans (they contain column names)
		if strings.Contains(id, "_ctl00_") || text == "Tarih" || text == "Durumu" || text == "İşlem Yeri" {
			return
		}

		// Validate the ID format
		parts := strings.Split(id, "_")
		if len(parts) < 3 {
			return
		}

		// Process based on the label type
		ts.processHTMLElement(id, text, processedDates, trackingData)
	})
}

// processHTMLElement handles different types of tracking data elements
func (ts *TrackingService) processHTMLElement(id, text string, processedDates map[string]bool, trackingData *[]models.Tracking) {
	// Identify what type of data this is based on the Label number
	if strings.Contains(id, "Label5") {
		// This is a date field, create a new entry if we haven't seen this date before
		if !processedDates[text] {
			processedDates[text] = true
			*trackingData = append(*trackingData, models.Tracking{Date: text})
			ts.Logger.Printf("Found date: %s", text)
		}
	} else if strings.Contains(id, "Label24") && len(*trackingData) > 0 {
		// This is a state field, find the corresponding tracking entry
		ts.updateTrackingField(trackingData, text, "state")
	} else if strings.Contains(id, "Label25") && len(*trackingData) > 0 {
		// This is a place field, find the corresponding tracking entry
		ts.updateTrackingField(trackingData, text, "place")
	}
}

// updateTrackingField updates a specific field in the most recent tracking entry that has an empty field
func (ts *TrackingService) updateTrackingField(trackingData *[]models.Tracking, value, fieldType string) {
	for i := len(*trackingData) - 1; i >= 0; i-- {
		switch fieldType {
		case "state":
			if (*trackingData)[i].State == "" {
				(*trackingData)[i].State = value
				ts.Logger.Printf("Found state: %s", value)
				return
			}
		case "place":
			if (*trackingData)[i].Place == "" {
				(*trackingData)[i].Place = value
				ts.Logger.Printf("Found place: %s", value)
				return
			}
		}
	}
}

// processTrackingResults handles the results of a tracking request
func (ts *TrackingService) processTrackingResults(trackingData []models.Tracking) {
	// Log if no tracking data found
	if len(trackingData) == 0 {
		ts.Logger.Println("Takip bilgisi bulunamadı! Lütfen takip numarasını veya bağlantınızı kontrol edin.")
		return
	}

	// Check if there are new tracking events
	currentCount := len(trackingData)
	if ts.PreviousTrackCount > 0 && currentCount > ts.PreviousTrackCount {
		ts.Logger.Printf("Yeni takip olayları tespit edildi! Sayı %d'den %d'e yükseldi",
			ts.PreviousTrackCount, currentCount)
		notification.NotifyNewStatus()
	}

	// Update previous count for next check
	ts.PreviousTrackCount = currentCount
}

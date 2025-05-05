package ui

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
	"upsTrack/internal/models"

	"github.com/charmbracelet/lipgloss"
)

// RenderTrackingEntry formats a tracking entry for display
func RenderTrackingEntry(track models.Tracking, styles models.TrackingStyles, isLatest bool) string {
	var dateInfo, stateInfo, placeInfo string

	if isLatest {
		// Special styling for the latest entry
		dateInfo = fmt.Sprintf("%s %s",
			styles.FirstDateLabel.Render("Tarih:"),
			styles.FirstDateValue.Render(track.Date))

		stateInfo = fmt.Sprintf("%s %s",
			styles.FirstStateLabel.Render("Durum:"),
			styles.FirstStateValue.Render(track.State))

		placeInfo = fmt.Sprintf("%s %s",
			styles.FirstPlaceLabel.Render("Konum:"),
			styles.FirstPlaceValue.Render(track.Place))
	} else {
		// Normal styling for older entries
		dateInfo = fmt.Sprintf("%s %s",
			styles.DateLabel.Render("Tarih:"),
			styles.DateValue.Render(track.Date))

		stateInfo = fmt.Sprintf("%s %s",
			styles.StateLabel.Render("Durum:"),
			styles.StateValue.Render(track.State))

		placeInfo = fmt.Sprintf("%s %s",
			styles.PlaceLabel.Render("Konum:"),
			styles.PlaceValue.Render(track.Place))
	}

	return lipgloss.JoinVertical(lipgloss.Left, dateInfo, stateInfo, placeInfo)
}

// PrintTrackingData displays the tracking information in a formatted way
func PrintTrackingData(trackingData []models.Tracking, styles models.TrackingStyles) {
	if len(trackingData) == 0 {
		noDataMsg := styles.NoDataMsg.Render("Takip bilgisi mevcut değil.")
		fmt.Println(noDataMsg)
		return
	}

	// Print header
	PrintHeader(styles)

	// Print each tracking entry with style
	for i, track := range trackingData {
		if i == 0 {
			// Special handling for the first (most recent) entry
			PrintLatestEntry(track, styles)
		} else {
			// Normal handling for older entries
			content := RenderTrackingEntry(track, styles, false)
			fmt.Println(content)
			fmt.Println(styles.Divider.Render(strings.Repeat("─", 50)))
		}
	}
}

// PrintHeader displays the title and timestamp
func PrintHeader(styles models.TrackingStyles) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	title := styles.Title.Render("UPS Kargo Takip Bilgileri")
	timestamp := styles.TimestampStyle.Render(fmt.Sprintf("%s itibariyle", currentTime))
	divider := styles.Divider.Render(strings.Repeat("─", 50))

	fmt.Println(title)
	fmt.Println(timestamp)
	fmt.Println(divider)
}

// PrintLatestEntry displays the most recent tracking entry with special formatting
func PrintLatestEntry(track models.Tracking, styles models.TrackingStyles) {
	content := RenderTrackingEntry(track, styles, true)

	// Wrap it in a special box
	firstEntryBox := styles.FirstEntryBox.Render(content)
	fmt.Println(lipgloss.NewStyle().Width(52).Align(lipgloss.Center).Render(" SON DURUM "))
	fmt.Println(firstEntryBox)
}

// PrintPrompt displays the tracking ID prompt to the user
func PrintPrompt() string {
	// Create a prompt style
	promptStyle := GetPromptStyle()

	// Display the prompt
	fmt.Println(promptStyle.Render("UPS Takip Numarasını Girin:"))

	// Read user input
	var trackingID string
	fmt.Scanln(&trackingID)

	// Trim whitespace
	return strings.TrimSpace(trackingID)
}

// ClearConsole clears the terminal screen
func ClearConsole() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

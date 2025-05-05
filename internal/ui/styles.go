package ui

import (
	"github.com/charmbracelet/lipgloss"
	"upsTrack/internal/models"
)

// TrackingStyles contains styled components for display

// NewTrackingStyles initializes and returns all the lipgloss styles used in the application
func NewTrackingStyles() models.TrackingStyles {
	return models.TrackingStyles{
		Title: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FF8A65")).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FF5722")).
			PaddingLeft(2).
			PaddingRight(2).
			MarginBottom(1),

		Divider: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#BDBDBD")).
			Width(50),

		DateLabel: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#4FC3F7")).
			Bold(true).
			PaddingRight(1),

		StateLabel: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#4FC3F7")).
			Bold(true).
			PaddingRight(1),

		PlaceLabel: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#4FC3F7")).
			Bold(true).
			PaddingRight(1),

		DateValue: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")),

		StateValue: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF8A65")),

		PlaceValue: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")),

		NoDataMsg: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#F44336")).
			Bold(true),

		TimestampStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#9E9E9E")).
			Italic(true),

		// Special styles for the first (latest) tracking entry
		FirstEntryBox: lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#4CAF50")).
			Padding(1).
			MarginTop(1).
			MarginBottom(1).
			Background(lipgloss.Color("#1E1E1E")),

		FirstDateLabel: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00E676")).
			Bold(true).
			PaddingRight(1),

		FirstStateLabel: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00E676")).
			Bold(true).
			PaddingRight(1),

		FirstPlaceLabel: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00E676")).
			Bold(true).
			PaddingRight(1),

		FirstDateValue: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Bold(true),

		FirstStateValue: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFEB3B")).
			Bold(true),

		FirstPlaceValue: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Bold(true),
	}
}

// GetPromptStyle creates a styled prompt for user input
func GetPromptStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		Padding(0, 1).
		Bold(true)
}

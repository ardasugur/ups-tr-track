package models

import (
	"github.com/charmbracelet/lipgloss"
)

// Tracking represents a single tracking event for a package
type Tracking struct {
	Date  string
	State string
	Place string
}

type TrackingStyles struct {
	Title          lipgloss.Style
	Divider        lipgloss.Style
	DateLabel      lipgloss.Style
	StateLabel     lipgloss.Style
	PlaceLabel     lipgloss.Style
	DateValue      lipgloss.Style
	StateValue     lipgloss.Style
	PlaceValue     lipgloss.Style
	NoDataMsg      lipgloss.Style
	TimestampStyle lipgloss.Style

	// Special styles for the first (latest) tracking entry
	FirstEntryBox   lipgloss.Style
	FirstDateLabel  lipgloss.Style
	FirstStateLabel lipgloss.Style
	FirstPlaceLabel lipgloss.Style
	FirstDateValue  lipgloss.Style
	FirstStateValue lipgloss.Style
	FirstPlaceValue lipgloss.Style
}

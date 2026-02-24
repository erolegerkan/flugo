package styles

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var verboseStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FFB366"))

var warningStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FFFF00"))

var errorStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FF0000"))

var successStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#33CC33"))

func RenderWithStyle(sentence string, renderType string) {
	switch renderType {
	case "verbose":
		fmt.Println(verboseStyle.Render(sentence))
	case "warning":
		fmt.Println(warningStyle.Render(sentence))
	case "error":
		fmt.Println(errorStyle.Render(sentence))
	case "success":
		fmt.Println(successStyle.Render(sentence))
	}
}

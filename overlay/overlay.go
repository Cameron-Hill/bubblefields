package overlay

import (
	"fmt"
	"strings"

	"github.com/Cameron-Hill/bubbleform/ansi"
	"github.com/charmbracelet/lipgloss"
)

func bigger(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func Overlay(base, overlay string) string {
	overlayWidth := lipgloss.Width(overlay)
	overlayHeight := lipgloss.Height(overlay)
	// overlayHeight += overlayHeight % 2
	width := bigger(lipgloss.Width(base), overlayWidth)
	height := bigger(lipgloss.Height(base), overlayHeight)

	// overlayWidth += (overlayWidth%2 + width%2) % 2

	left := (width - overlayWidth) / 2
	top := (height - overlayHeight) / 2
	right := left + overlayWidth
	bottom := top + overlayHeight

	renderer := lipgloss.DefaultRenderer()

	// base = Strip(base)

	content := strings.Split(renderer.Place(width, height, lipgloss.Left, lipgloss.Top, base), "\n")
	// overlayPad := overlayHeight - lipgloss.Height(overlay)
	// overlay += strings.Repeat("\n", overlayPad)
	overlayLines := strings.Split(overlay, "\n")

	for i := top; i < bottom; i++ {
		var with string
		if len(overlayLines) > i-top {
			with = fmt.Sprintf("%-*s", overlayWidth, overlayLines[i-top]) // Pad with spaces to the right
		} else {
			with = strings.Repeat(" ", overlayWidth) // Pad with spaces at the bottom
		}

		// Handle styles
		// 1. Get actual index at left and right
		actualLeft, lerr := ansi.ActualIndex(content[i], left)
		actualRight, rerr := ansi.ActualIndex(content[i], right)

		if lerr != nil || rerr != nil {
			continue
		}

		// 2. Determine what the style should be at overlay termination
		terminationStyles := strings.Join(ansi.ActiveANSICodes(content[i], right), "")

		// 3. Cancel any existing styles where the overlay starts
		with = "\x1b[0m" + with

		content[i] = content[i][:actualLeft] + with + terminationStyles + content[i][actualRight:]

	}

	return strings.Join(content, "\n")
}

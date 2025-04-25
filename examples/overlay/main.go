package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Cameron-Hill/bubbleform/overlay"
	"github.com/charmbracelet/lipgloss"
)

const (
	DEFAULT_WIDTH          = 20
	DEFAULT_HEIGHT         = 10
	DEFAULT_OVERLAY_WIDTH  = 8
	DEFAULT_OVERLAY_HEIGHT = 4
)

func getOverlay(width, height int) string {
	content := strings.TrimSuffix(strings.Repeat(strings.Repeat("o", width)+"\n", height), "\n")

	return lipgloss.NewStyle().Border(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("10")).Foreground(lipgloss.Color("6")).Render(content)
}

func getArgs() (contentWidth, contentHeight, overlayWidth, overlayHeight int) {
	contentWidth = DEFAULT_WIDTH
	contentHeight = DEFAULT_HEIGHT
	overlayWidth = DEFAULT_OVERLAY_WIDTH
	overlayHeight = DEFAULT_OVERLAY_HEIGHT

	for i := 1; i < 5; i++ {
		if len(os.Args) > i {
			arg, err := strconv.Atoi(os.Args[i])
			if err == nil {
				switch i {
				case 1:
					contentWidth = arg
				case 2:
					contentHeight = arg
				case 3:
					overlayWidth = arg
				case 4:
					overlayHeight = arg
				}
			} else {
				fmt.Printf("Error: Invalid argument %d\n", i)
			}
		} else {
			break
		}
	}
	return
}

func main() {
	width, height, overlayWidth, overlayHeight := getArgs()
	builder := strings.Builder{}
	for i := 0; i < height; i++ {
		if i != 4 {
			builder.WriteString(strings.Repeat("#", width) + "\n")
		} else {
			builder.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("21")).Render(strings.Repeat("#", width)) + "\n")
		}
	}

	base := strings.TrimSuffix(builder.String(), "\n")
	fmt.Println("============== BASE ===============")
	fmt.Println(base + "\n")
	fmt.Println("============= OVERLAY =============")
	overlayStr := getOverlay(overlayWidth, overlayHeight)
	fmt.Println(overlayStr + "\n")
	output := overlay.Overlay(base, overlayStr)
	fmt.Println("============== OUTPUT ==============")
	fmt.Println(output)
}

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func getOverlay(width, height int) string {
	content := strings.TrimSuffix(strings.Repeat(strings.Repeat("o", width)+"\n", height), "\n")

	return lipgloss.NewStyle().Border(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("10")).Foreground(lipgloss.Color("6")).Render(content)
}

func getArgs() (contentWidth, contentHeight, overlayWidth, overlayHeight int) {
	contentWidth = 20
	contentHeight = 10
	overlayWidth = 8
	overlayHeight = 4

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
	if true {
		model := InitialModel()
		p := tea.NewProgram(
			model,
			tea.WithAltScreen(),
		)
		_, err := p.Run()
		if err != nil {
			fmt.Print("lkasjglkadfjgl;sdkfjg")
		}

	} else {
		width, height, overlayWidth, overlayHeight := getArgs()
		builder := strings.Builder{}
		for i := range height {
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
		output := Overlay(base, overlayStr)
		fmt.Println("============== OUTPUT ==============")
		fmt.Println(output)
	}
}

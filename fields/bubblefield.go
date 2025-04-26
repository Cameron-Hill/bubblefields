package fields

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type BubbleField struct {
	model    tea.Model
	getValue func(tea.Model) string
	err      error
}

func (f *BubbleField) Value() string {
	return f.getValue(f.model)
}

func (f *BubbleField) SetError(err error) {
	f.err = err
}

func (f *BubbleField) Init() tea.Cmd {
	return nil
}

func (f *BubbleField) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	model, cmd := f.model.Update(msg)
	f.model = model
	return f, cmd
}

func (f *BubbleField) View() string {
	return lipgloss.JoinVertical(lipgloss.Bottom,
		f.model.View(),
		lipgloss.NewStyle().Height(1).Foreground(lipgloss.Color("red")).Render(f.err.Error()),
	)
}

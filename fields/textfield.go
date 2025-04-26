package fields

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TextField struct {
	textInput  textinput.Model
	label      string
	err        error
	style      lipgloss.Style
	errorStyle lipgloss.Style
}

type TextFieldOption func(*TextField)

func NewTextField(label string, opts ...TextFieldOption) TextField {
	ti := textinput.New()

	tf := TextField{
		textInput:  ti,
		label:      label,
		style:      lipgloss.NewStyle(),
		errorStyle: lipgloss.NewStyle().Height(1).Foreground(lipgloss.Color("3")),
	}

	for _, opt := range opts {
		opt(&tf)
	}

	return tf
}

// Option to set the style of the text field
func WithStyle(style lipgloss.Style) TextFieldOption {
	return func(tf *TextField) {
		tf.style = style
	}
}

// Option to set the error style
func WithErrorStyle(style lipgloss.Style) TextFieldOption {
	return func(tf *TextField) {
		tf.errorStyle = style
	}
}

// Option to set the placeholder
func WithPlaceholder(placeholder string) TextFieldOption {
	return func(tf *TextField) {
		tf.textInput.Placeholder = placeholder
	}
}

// Option to set the initial value
func WithInitialValue(val string) TextFieldOption {
	return func(tf *TextField) {
		tf.textInput.SetValue(val)
	}
}

// Option to set focus
func WithFocus(focus bool) TextFieldOption {
	return func(tf *TextField) {
		if focus {
			tf.textInput.Focus()
		} else {
			tf.textInput.Blur()
		}
	}
}

func (f *TextField) Focus() {
	f.textInput.Focus()
}

func (f *TextField) Blur() {
	f.textInput.Blur()
}

func (f *TextField) SetValue(val string) {
	f.textInput.SetValue(val)
}

func (f *TextField) Value() string {
	return f.textInput.Value()
}

func (f *TextField) SetError(err error) {
	f.err = err
}

func (f *TextField) ClearError() {
	f.err = nil
}

func (f *TextField) Init() tea.Cmd {
	return textinput.Blink
}

func (f *TextField) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	f.textInput, cmd = f.textInput.Update(msg)
	return f, cmd
}

func (f *TextField) View() string {
	output := fmt.Sprintf("%s\n%s", f.label, f.textInput.View())

	if f.err != nil {
		output = lipgloss.JoinVertical(lipgloss.Left,
			f.style.Render(output),
			f.errorStyle.Render(f.err.Error()),
		)
	}

	return f.style.Render(output)
}

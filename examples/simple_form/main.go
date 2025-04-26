package main

import (
	"fmt"
	"log"

	"github.com/Cameron-Hill/bubbleform"
	"github.com/Cameron-Hill/bubbleform/fields"
	tea "github.com/charmbracelet/bubbletea"
)

// User data structure with validation tags
type User struct {
	Name     string `validate:"required,min=2"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

type model struct {
	form          *bubbleform.Form
	nameField     *fields.TextField
	emailField    *fields.TextField
	passwordField *fields.TextField
	currentField  int
	submitted     bool
	user          User
}

func initialModel() model {
	user := User{}
	form := bubbleform.NewForm(&user)

	nameField := fields.NewTextField("Name")
	emailField := fields.NewTextField("Email")
	passwordField := fields.NewTextField("Password")

	// Bind fields to form
	if err := form.Bind("Name", &nameField); err != nil {
		log.Fatal(err)
	}
	if err := form.Bind("Email", &emailField); err != nil {
		log.Fatal(err)
	}
	if err := form.Bind("Password", &passwordField); err != nil {
		log.Fatal(err)
	}

	nameField.Focus()

	return model{
		form:          form,
		nameField:     &nameField,
		emailField:    &emailField,
		passwordField: &passwordField,
		currentField:  0,
		user:          user,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		m.nameField.Init(),
		m.emailField.Init(),
		m.passwordField.Init(),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.currentField < 2 {
				// Move to next field
				m.currentField++
				if m.currentField == 0 {
					m.nameField.Focus()
					m.emailField.Blur()
					m.passwordField.Blur()
				} else if m.currentField == 1 {
					m.nameField.Blur()
					m.emailField.Focus()
					m.passwordField.Blur()
				} else {
					m.nameField.Blur()
					m.emailField.Blur()
					m.passwordField.Focus()
				}
			} else {
				// Submit form
				m.submitted = true
				if m.form.Submit() {
					return m, tea.Quit
				}
			}
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	}

	// Update current field
	var cmd tea.Cmd
	if m.currentField == 0 {
		_, cmd = m.nameField.Update(msg)
	} else if m.currentField == 1 {
		_, cmd = m.emailField.Update(msg)
	} else {
		_, cmd = m.passwordField.Update(msg)
	}
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	if m.submitted && len(m.form.GetErrors()) == 0 {
		return fmt.Sprintf("Form submitted successfully!\nName: %s\nEmail: %s\nPassword: (hidden)",
			m.user.Name, m.user.Email)
	}

	return fmt.Sprintf(
		"%s\n\n%s\n\n%s\n\n(Enter to advance, Esc to quit)",
		m.nameField.View(),
		m.emailField.View(),
		m.passwordField.View(),
	)
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

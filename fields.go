package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Field struct {
	model      tea.Model
	key        string
	editable   bool
	selectable bool
}

type FieldOpt func(f *Field)

func (f Field) New(model Model, opts ...FieldOpt) Field {
	field := Field{
		model: model,
	}
	for _, opt := range opts {
		opt(&field)
	}

	return field
}

func EditableField(f *Field) {
	f.editable = true
}

func SelectableField(f *Field) {
	f.selectable = true
}

type Model struct {
	selected int
	fields   []Field
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			m.selected--
		case "down":
			m.selected++
		case "ctrl+c", "ctrl+d", "q", "esc":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	return "wooo"
}

func InitialModel() Model {
	return Model{
		selected: 0,
	}
}

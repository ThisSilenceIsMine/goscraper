package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

type model struct {
	links textarea.Model
	err   error
}

func initialModel() model {
	ti := textarea.New()
	ti.Placeholder = "https://guide.elm-lang.org/architecture/"
	ti.Focus()

	return model{
		links: ti,
		err:   nil,
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			if m.links.Focused() {
				m.links.Blur()
			}
		case tea.KeyCtrlC:
			return m, tea.Quit
		default:
			if !m.links.Focused() {
				cmd = m.links.Focus()
				cmds = append(cmds, cmd)
			}
		}

	// We handle errors just like any other message
	case error:
		m.err = msg
		return m, nil
	}

	m.links, cmd = m.links.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	return fmt.Sprintf(
		"Enter your links: \n\n%s\n\n%s",
		m.links.View(),
		"(ctrl+c to quit)",
	) + "\n\n"
}

func main() {
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

package frontend

import (
	"TerminalAI/ai"
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg error
)

type model struct {
	textInput textinput.Model
	err       error
}

func InitialModel() model {
	ti := textinput.New()
	ti.Placeholder = "type here"
	ti.Focus()
	ti.CharLimit = 3000
	ti.Width = 3000
	ti.Prompt = "user > "

	return model{
		textInput: ti,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.Type {

		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyEnter:
			if m.textInput.Value() == "" {
				return m, nil
			} else if m.textInput.Value()[0] == '/' {
				if m.textInput.Value() == "/clear" {
					fmt.Print("\033[H\033[2J")
					m.textInput.Reset()
					return m, nil
				} else if m.textInput.Value() == "/exit" || m.textInput.Value() == "/quit" || m.textInput.Value() == "/q" || m.textInput.Value() == "/bye" { 
					return m, tea.Quit
				} else if m.textInput.Value() == "/help" {
					m.textInput.Reset()
					fmt.Println("Available commands:",
					"\n  - /clear: Clear the screen",
					"\n  - /exit or /quit or /bye or /q: Exit the program",
					"\n  - /help: shows commands",
					"\n",
					)
					return m, nil
				}
			} else {
				response := ai.MakeRequest(m.textInput.Value())
				fmt.Println("\n;) > ", response, "\n")
				m.textInput.Reset()
			}
	}

	case errMsg:
		m.err = msg
		return m, nil

	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s",
		// "(esc to quit)",
		m.textInput.View(),
	) + "\n"
}
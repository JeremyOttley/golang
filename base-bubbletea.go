package main

import (
        "fmt"
        "os"

        "github.com/charmbracelet/bubbles/key"
        "github.com/charmbracelet/bubbles/spinner"
        tea "github.com/charmbracelet/bubbletea"
        "github.com/charmbracelet/lipgloss"
)

func main() {
	t := textinput.New()
	t.Focus()

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	initialModel := Model{
		textInput: t,
		spinner:   s,
		typing:    true,
	}
	err := tea.NewProgram(initialModel, tea.WithAltScreen()).Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type Model struct {
	textInput textinput.Model
	spinner   spinner.Model
	typing bool
	loading bool
	quitting  bool
	err       error
}

var quitKeys = key.NewBinding(
        key.WithKeys("q", "esc", "ctrl+c"),
        key.WithHelp("", "press q to quit"),
)

var enter = key.NewBinding(
	key.WithKeys("enter"),
	key.WithKeys("", "press Enter to search"),
)

func (m Model) MyCustomFunction(input string) tea.Cmd {
	return func() tea.Msg {
	// Body
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyMsg:

		if key.Matches(msg, quitKeys){
			m.quitting = true
			return m, tea.Quit
		}
		return m, nil

		if key.Matches(msg, enter){
			if m.typing {

				userInput := strings.TrimSpace(m.textInput.Value())

				if userInput != "" { //guard clause in case of blank string

					m.typing = false
					m.loading = true

					return m, tea.Batch(
						spinner.Tick,
						m.MyCustomFunction(userInput),
					)
				}
			}
		}
		return m, nil

	case m.err:
		m.err = msg
		return m, nil


	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}





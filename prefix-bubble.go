package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/antchfx/xmlquery"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
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
	typing    bool
	loading   bool
	err       error
	publisher string
}

type GotPublisher struct {
	Err       error
	Publisher string
}

func (m Model) fetchPublisher(prefix string) tea.Cmd {
	return func() tea.Msg {
		crossrefApi := fmt.Sprintf("https://doi.crossref.org/getPrefixPublisher/?prefix=%s", prefix)

		doc, err := xmlquery.LoadURL(crossrefApi)
		if err != nil {
			return GotPublisher{Err: err}
		}

		data := xmlquery.FindOne(doc, "//xml/publisher")

		publisher := data.SelectElement("publisher_name").InnerText()

		return GotPublisher{Publisher: publisher}
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c":

			return m, tea.Quit

		case "enter":

			if m.typing {

				prefix := strings.TrimSpace(m.textInput.Value())

				if prefix != "" { //guard clause in case of blank string

					m.typing = false
					m.loading = true

					return m, tea.Batch(
						spinner.Tick,
						m.fetchPublisher(prefix),
					)
				}
			}

		case "esc":

			if !m.typing && !m.loading {
				m.typing = true
				m.err = nil
				return m, nil
			}
		}

	case GotPublisher:

		m.loading = false

		if err := msg.Err; err != nil {
			m.err = err
			return m, nil
		}

		m.publisher = msg.Publisher

	default:
		return m, nil

	}

	if m.typing {
		var cmd tea.Cmd
		m.textInput, cmd = m.textInput.Update(msg)
		return m, cmd
	}

	if m.loading {
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m Model) View() string {

	if m.typing {
		return fmt.Sprintf("Enter a DOI prefix:\n%s", m.textInput.View())
	}

	if m.loading {
		return fmt.Sprintf("%s fetching prefix... please wait.", m.spinner.View())
	}

	if err := m.err; err != nil {
		return fmt.Sprintf("Could not fetch prefix: %v", err)
	}

	return fmt.Sprintf("%s", m.publisher)

}

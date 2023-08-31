package main

import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
    "strings"

    "github.com/antchfx/xmlquery"
    tea "github.com/charmbracelet/bubbletea"
)

const url = "https://doi.crossref.org/getPrefixPublisher/?prefix=10.1215" //+prefix

type model struct {
    xml string
    err error
}

type xmlMsg string

type errMsg struct {
    error
}

func(e errMsg) Error() string {
    return e.error.Error()
}

func main() {
    p: = tea.NewProgram(model {})
    if _,
    err: = p.Run();err != nil {
        log.Fatal(err)
    }
}

func(m model) Init() tea.Cmd {
    return checkPrefix
}

func(m model) Update(msg tea.Msg)(tea.Model, tea.Cmd) {
    switch msg: = msg.(type) {
        case tea.KeyMsg:
            switch msg.String() {
                case "q", "ctrl+c", "esc":
                    return m, tea.Quit
                default:
                    return m, nil
            }

        case xmlMsg:
            m.xml = string(msg)
            return m, tea.Quit

        case errMsg:
            m.err = msg
            return m, nil

        default:
            return m, nil
    }
}

func(m model) View() string {
    s: = fmt.Sprintf("Checking CrossRef...\n\n")
    if m.err != nil {
        s += fmt.Sprintf("something went wrong: %s", m.err)
    } else if m.xml != "" {
        s += fmt.Sprintf("%s", m.xml)
    }
    return s + "\n"
}

func checkPrefix() tea.Msg {
    res, err: = http.Get(url)
    if err != nil {
        return errMsg {
            err
        }
    }

    body, err: = ioutil.ReadAll(res.Body)
    if err != nil {
        return errMsg {
            err
        }
    }
    string_body: = string(body)
    doc, err: = xmlquery.Parse(strings.NewReader(string_body))
    if err != nil {
        panic(err)
    }
    data := xmlquery.FindOne(doc, "//xml/publisher")

    return xmlMsg(data.SelectElement("publisher_name").InnerText())

}

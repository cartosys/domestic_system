package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "github.com/cartosys/domestic_system/libs/wallet"
    tea "github.com/charmbracelet/bubbletea"
    "os"
)

type model struct {
    choices  []string           // items on the to-do list
    cursor   int                // which to-do list item our cursor is pointing at
    selected map[int]struct{}   // which to-do items are selected
}
func initialModel(list []string) model {
	return model{
		// Our to-do list is a grocery list
    choices:  list,
		//choices:  inputs,

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
    // Just return `nil`, which means "no I/O right now, please."
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    // Is it a key press?
    case tea.KeyMsg:

        // Cool, what was the actual key pressed?
        switch msg.String() {

        // These keys should exit the program.
        case "ctrl+c", "q":
            return m, tea.Quit

        // The "up" and "k" keys move the cursor up
        case "up", "k":
            if m.cursor > 0 {
                m.cursor--
            }

        // The "down" and "j" keys move the cursor down
        case "down", "j":
            if m.cursor < len(m.choices)-1 {
                m.cursor++
            }

        // The "enter" key and the spacebar (a literal space) toggle
        // the selected state for the item that the cursor is pointing at.
        case "enter", " ":
            _, ok := m.selected[m.cursor]
            if ok {
                delete(m.selected, m.cursor)
            } else {
                m.selected[m.cursor] = struct{}{}
            }
        }
    }

    // Return the updated model to the Bubble Tea runtime for processing.
    // Note that we're not returning a command.
    return m, nil
}
func (m model) View() string {
    // The header
    s := "Which account should we use?\n\n"

    // Iterate over our choices
    for i, choice := range m.choices {

        // Is the cursor pointing at this choice?
        cursor := " " // no cursor
        if m.cursor == i {
            cursor = ">" // cursor!
        }

        // Is this choice selected?
        checked := " " // not selected
        if _, ok := m.selected[i]; ok {
            checked = "x" // selected!
        }

        // Render the row
        s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
    }

    // The footer
    s += "\nPress q to quit.\n"

    // Send the UI for rendering
    return s
}
func main() {

    fmt.Println(wallet.Main())

    /*values := map[string]string{"jsonrpc":"2.0","method":"eth_accounts","id":"1"}
    json_data, err := json.Marshal(values)

    if err != nil {
        log.Fatal(err)
    }

    resp, err := http.Post("http://127.0.0.1:8545", "application/json",
        bytes.NewBuffer(json_data))

    if err != nil {
        log.Fatal(err)
    }

    var res map[string]interface{}

    json.NewDecoder(resp.Body).Decode(&res)

    aInterface := res["result"].([]interface{})
    aString := make([]string, len(aInterface))
    for i, v := range aInterface {
        aString[i] = v.(string)
    }

    p := tea.NewProgram(initialModel(aString))
    if err := p.Start(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }*/
}

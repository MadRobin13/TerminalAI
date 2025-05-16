package main

import (
	"TerminalAI/ai"
	"TerminalAI/frontend"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	ai.SetUpAPI()

	// fmt.Println(ai.MakeRequest("How are you"))

	// Initialize the model
	p := tea.NewProgram(frontend.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}
}
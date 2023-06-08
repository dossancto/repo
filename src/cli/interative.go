package cli

import (
	"fmt"
	"log"
	"strings"

	"github.com/lu-css/repo/src/apis"
	"github.com/manifoldco/promptui"
)

func startInterative() {
	cmd := getCommand()

	println(cmd)
}

func getCommand() string {
	commands := listCommands()

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U000027A1 {{ . | cyan }}",
		Inactive: "  {{ . | cyan }} ",
		Selected: "Commands: {{ . | red | cyan }}",
	}

	prompt := promptui.Select{
		Label:     "Select a Model",
		Items:     commands,
		Templates: templates,
		Size:      15,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return strings.ToLower(commands[i])
}

func listCommands() []string {
	commands := []string{
		"clone",
	}
	return commands
}

func chooseRemote() string {
	remotes, err := apis.ListRemotes()

	if len(remotes) == 1 {
		return remotes[0]
	}

	if err != nil {
		log.Fatal(err)
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U000027A1 {{ . | cyan }}",
		Inactive: "  {{ . | cyan }} ",
		Selected: "Remotes: {{ . | red | cyan }}",
	}

	prompt := promptui.Select{
		Label:     "Choose a remote",
		Items:     remotes,
		Templates: templates,
		Size:      5,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return strings.ToLower(remotes[i])
}

package commands

import (
	"fmt"
	"log"
	"strings"

	"github.com/lu-css/repo/src/apis"
	"github.com/manifoldco/promptui"
)

/*
If userName contains a "/" automaticly clone repo.
Otherwise guess witch repository the user wants to clone.
*/
func CloneCliRepo() {
	userName := getGithubUser()
	repo := getUserRepos(userName)

	CloneRepo(userName, repo)

}

func CloneRepo(user string, repo string) {
	out, err := apis.Clone(user, repo)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)
}

func getGithubUser() string {
	prompt := promptui.Prompt{
		Label: "Github user",
	}

	result, err := prompt.Run()

	if err != nil {
		log.Fatal(err)
		return ""
	}

	return result
}

func getUserRepos(user string) string {
	repos, err := apis.GetAllPublicRepos(user)

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U000027A1 {{ . | cyan }}",
		Inactive: "  {{ . | cyan }} ",
		Selected: "Repositories: {{ . | red | cyan }}",
	}

	prompt := promptui.Select{
		Label:     "Select a Model",
		Items:     repos,
		Templates: templates,
		Size:      15,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return strings.ToLower(repos[i])

}

package cli

import (
	"fmt"
	"log"

	"github.com/lu-css/repo/src/apis"
	"github.com/lu-css/repo/src/commands"
	"github.com/lu-css/repo/src/utils"
)

func templateUrl(args []string) {
	user := utils.GetUser(args, 0)
	repo := args[1]

	url := apis.BuildRepoURL(user, repo)
	fmt.Print(url)
}

func cloneRepo(args []string) {
	var repo string
	user := utils.GetUser(args, 0)

	if len(args) <= 1 {
		repo = commands.GetUserRepos(user)
	} else {
		repo = args[1]
	}

	commands.CloneRepo(user, repo)
}

func getActualRepoUrl(args []string) {
	var remote string

	if len(args) < 1 {
		remote = chooseRemote()
	} else {
		remote = args[0]
	}

	url, err := apis.RemoteGetUrl(remote)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(url)
}

package utils

import (
	"log"

	"github.com/lu-css/repo/src/apis"
)

func GetUser(args []string, index int) string {
	var user string
	var err error

	if (len(args) - 1) >= index {
		user = args[index]
	}

	if user != "!" && user != "" {
		return user
	}

	user, err = apis.GetActualGitUser()

	if err != nil {
		log.Fatal(err)
	}

	return user
}

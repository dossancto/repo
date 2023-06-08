package apis

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

const BASE_URL = "https://api.github.com/"

/*
Retuns all public repositories from user.
*/
func GetAllPublicRepos(user string) ([]string, error) {
	resp, err := http.Get(BASE_URL + "users/" + user + "/repos")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	var repos Repositories

	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		log.Fatal(err)
	}

	return repos.RepoNames(), errors.New("Something went wrong")
}

func BuildRepoURL(user string, repo string) string {
	return "https://github.com/" + user + "/" + repo
}

package apis

import (
	"bufio"
	"errors"
	"os/exec"
	"strings"
)

/*
Clones the specified repository.

Returns `err` as nil if success.
*/
func Clone(user string, repo string) (string, error) {
	cmd := exec.Command("git", "clone", BuildRepoURL(user, repo))

	out, err := cmd.CombinedOutput()

	if err != nil {
		return "", err
	}

	return string(out), nil
}

func RemoteGetUrl(remote string) (string, error) {
	cmd := exec.Command("git", "remote", "get-url", remote)

	out, err := cmd.CombinedOutput()

	if err != nil {
		return "", err
	}

	return string(out), nil
}

func ListRemotes() ([]string, error) {
	cmd := exec.Command("git", "remote")

	out, err := cmd.CombinedOutput()

	if err != nil {
		return nil, err
	}

  remotes := linesToArray(string(out))

  if len(remotes) == 0 {
    return nil, errors.New("No one remote.")
  }

	return remotes, nil
}

func linesToArray(buffer string) []string {
	var lines []string

	reader := strings.NewReader(buffer)
	scan := bufio.NewScanner(reader)

	for scan.Scan() {
		line := scan.Text()
		lines = append(lines, line)
	}

	return lines
}

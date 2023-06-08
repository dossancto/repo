package apis

import (
	"os/exec"
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

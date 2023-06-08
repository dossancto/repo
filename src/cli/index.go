package cli

func WithArgs(args []string) {
	command := args[0]
	newArgs := args[1:]

	switch command {
	case ("t"):
		templateUrl(newArgs)
		break

	case ("c"):
		cloneRepo(newArgs)
		break

	case ("u"):
		getActualRepoUrl(newArgs)
		break
	}
}

func Interative() {
	startInterative()
}

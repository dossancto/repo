# Repo

## Build

`go build -o bin/`

## Usage:

`!` works as the configured user.name user.

Example:

`repo c ! some-repo` same as `repo c {your-profile} some-repo` 

### repo clone

Clone the selected project.

`repo c` - Uses the user.name git user, list the public repositories and clone the selected repository
`repo c {github user} {repository}` - Uses the user.name git user, list the public repositories and clone the selected repository

### repo template

Prints the repository URL to stdout

`repo t some-user some-repo`

### Repo remote

Prints the remote URL to the STDOUT

`repo u`

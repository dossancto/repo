package apis

type Repositories []struct {
  Name string `json:"Name"`

}

func (r Repositories) RepoNames() []string {
  var repos []string

  for _, repo := range r  {
    repos = append(repos, repo.Name)
  }

  return repos

}

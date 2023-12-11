package data

type Edge struct {
	Permission string
	Node       struct {
		Login string
	}
}

type RepoInfo struct {
	DatabaseId    int    `json:"databaseId"`
	Name          string `json:"name"`
	Visibility    string `json:"visibility"`
	Collaborators struct {
		Edges []Edge
	} `graphql:"collaborators(first:1, query: $user)"`
}

type OrganizationUserQuery struct {
	Organization struct {
		Repositories struct {
			Nodes    []RepoInfo
			PageInfo struct {
				EndCursor   string
				HasNextPage bool
			}
		} `graphql:"repositories(first: 100, after: $endCursor)"`
	} `graphql:"organization(login: $owner)"`
}

type RepoSingleQuery struct {
	Repository RepoInfo `graphql:"repository(owner: $owner, name: $name)"`
}

type RepoCollaborators struct {
	Login string `json:"login"`
	Id    int    `json:"id"`
	Type  string `json:"type"`
}

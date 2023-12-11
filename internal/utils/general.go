package utils

import (
	"fmt"
	"io"
	"log"

	"github.com/cli/go-gh/pkg/api"
	"github.com/katiem0/gh-collaborators/internal/data"
	"github.com/shurcooL/graphql"
	"go.uber.org/zap"
)

type Getter interface {
	GetGuestCollaborators(owner string) ([]byte, error)
	GetOrgRepositoryPermissions(owner string, user string, endCursor *string) (*data.OrganizationUserQuery, error)
}

type APIGetter struct {
	gqlClient  api.GQLClient
	restClient api.RESTClient
}

func NewAPIGetter(gqlClient api.GQLClient, restClient api.RESTClient) *APIGetter {
	return &APIGetter{
		gqlClient:  gqlClient,
		restClient: restClient,
	}
}

func (g *APIGetter) GetOrgGuestCollaborators(owner string) ([]byte, error) {
	url := fmt.Sprintf("orgs/%s/outside_collaborators", owner)
	zap.S().Debugf("Reading in repository collaborators from %v", url)
	resp, err := g.restClient.Request("GET", url, nil)
	if err != nil {
		log.Printf("Body read error, %v", err)
	}
	defer resp.Body.Close()
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Body read error, %v", err)
	}
	return responseData, err
}

func (g *APIGetter) GetOrgRepositoryPermissions(owner string, user string, endCursor *string) (*data.OrganizationUserQuery, error) {
	query := new(data.OrganizationUserQuery)
	variables := map[string]interface{}{
		"endCursor": (*graphql.String)(endCursor),
		"owner":     graphql.String(owner),
		"user":      graphql.String(user),
	}
	err := g.gqlClient.Query("getOrganizationRepoPermissions", &query, variables)

	return query, err
}

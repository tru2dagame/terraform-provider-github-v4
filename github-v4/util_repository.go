package github

import (
	"context"
	"github.com/shurcooL/githubv4"
)

const (
	CODEOWNERS_EXISTS        = "exists"
	REPOSITORY_COLLABORATORS = "collaborators"
	REPOSITORY_ID            = "repository_id"
	REPOSITORY_NAME          = "name"
)

func getRepositoryID(name string, meta interface{}) (githubv4.ID, error) {
	var query struct {
		Repository struct {
			ID githubv4.ID
		} `graphql:"repository(owner:$owner, name:$name)"`
	}
	variables := map[string]interface{}{
		"owner": githubv4.String(meta.(*Organization).Name),
		"name":  githubv4.String(name),
	}
	ctx := context.Background()
	client := meta.(Organization).Client
	err := client.Query(ctx, &query, variables)
	if err != nil {
		return nil, err
	}

	return query.Repository.ID, nil
}

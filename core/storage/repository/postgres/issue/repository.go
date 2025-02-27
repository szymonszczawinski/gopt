package issue

import (
	"context"
	"errors"
	"gopt/core/domain/issue"
	"gopt/core/storage/repository/postgres"
	"gopt/coreapi"
	"log/slog"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

var ErrIssueNotFound = errors.New("issue not found")

const (
	ISSUE_SELECT_BY_KEY = "SELECT i.id, i.created, i.updated, i.name, i.item_key,  " +
		" lcs.name as state_name, lc.name as itemType," +
		" CONCAT(c.last_name,', ',c.first_name) as creator_name, " +
		" CONCAT(a.last_name,', ',a.first_name) as assignee_name, " +
		" i.project_id, i.project_key " +
		" from issue i " +
		" JOIN lifecyclestate lcs ON lcs.id = i.state_id" +
		" JOIN lifecycle lc ON lc.id = i.lifecycle_id" +
		" JOIN users c ON c.id = i.created_by_id" +
		" JOIN users a ON a.id = i.assigned_to_id" +
		" WHERE i.item_key = $1"
)

type issueRepositoryPostgres struct {
	lockDb *sync.RWMutex
	db     postgres.IPostgresDatabase

	eg  *errgroup.Group
	ctx context.Context
}

func NewIssueRepositoryPostgres(eg *errgroup.Group, ctx context.Context, db postgres.IPostgresDatabase) *issueRepositoryPostgres {
	instance := issueRepositoryPostgres{
		lockDb: &sync.RWMutex{},
		db:     db,
		eg:     eg,
		ctx:    ctx,
	}
	return &instance
}

func (repo *issueRepositoryPostgres) StartComponent() {
}

func (repo issueRepositoryPostgres) GetIssue(issueKey string) coreapi.Result[issue.Issue] {
	result := repo.db.NewSelectOne(ISSUE_SELECT_BY_KEY, issueKey)
	if result == nil {
		return coreapi.NewResult[issue.Issue](issue.Issue{}, ErrIssueNotFound)
	}
	var row struct {
		created, updated                                              time.Time
		itemKey, name, state, creator, assingee, itemType, projectKey string
		id, projectId                                                 int
	}

	err := result.Scan(&row.id, &row.created, &row.updated, &row.name, &row.itemKey, &row.state, &row.itemType,
		&row.creator, &row.assingee, &row.projectId, &row.projectKey)
	if err != nil {
		slog.Error("repo get issue", "err", err, "row", row)
	}
	issue := issue.NewIssueFromRepo(row.id, row.itemKey, row.name, row.itemType, row.projectId, row.projectKey)
	return coreapi.NewResult(issue, nil)
}

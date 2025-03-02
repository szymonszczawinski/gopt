package issue

import "gopt/coreapi"

type IIssueRepo interface {
	GetIssue(key string) coreapi.Result[Issue]
}
type issueService struct {
	repo IIssueRepo
}

func NewIssueService(repo IIssueRepo) *issueService {
	return &issueService{
		repo: repo,
	}
}

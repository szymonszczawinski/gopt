package issue

import "errors"

type CreateIssue struct {
	IssueType  string `json:"issue_type"`
	Name       string `json:"name"`
	ProjectKey string `json:"project_key"`
}

func NewCreateIssue(issueType, name, projectKey string) (CreateIssue, error) {
	command := CreateIssue{
		IssueType:  issueType,
		Name:       name,
		ProjectKey: projectKey,
	}

	return command, command.validate()
}

func (c CreateIssue) validate() error {
	result := ""
	if len(c.Name) == 0 {
		result = "Name must not be empty;"
	}
	if len(c.ProjectKey) == 0 {
		result += "Parent project must not be empty;"
	}
	if len(c.IssueType) == 0 {
		result += "Issue type must not be empty;"
	}
	if len(result) != 0 {
		return errors.New(result)
	} else {
		return nil
	}
}

type GetIssue struct {
	IssueKey string `json:"issue_key"`
}

func NewGetIssue(issueKey string) (GetIssue, error) {
	command := GetIssue{
		IssueKey: issueKey,
	}

	return command, command.validate()
}

func (c GetIssue) validate() error {
	result := ""
	if len(c.IssueKey) == 0 {
		result = "Issue key must not be empty;"
	}
	if len(result) != 0 {
		return errors.New(result)
	} else {
		return nil
	}
}

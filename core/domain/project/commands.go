package project

import "errors"

// Command for creating new project
type CreateProject struct {
	ProjectKey string `json:"project_key"`
	Name       string `json:"project_name"`
}

func NewCreateProject(projectKey, projectName string) (CreateProject, error) {
	command := CreateProject{
		ProjectKey: projectKey,
		Name:       projectName,
	}
	return command, command.validate()
}

func (c CreateProject) validate() error {
	var result string
	if len(c.Name) == 0 {
		result = "Name must not be empty.\n"
	}
	if len(c.ProjectKey) == 0 {
		result += "Key must not be empty"
	}
	if len(result) != 0 {
		return errors.New(result)
	} else {
		return nil
	}
}

// Command for get single project by its key
type GetProject struct {
	ProjectKey string `json:"project_key"`
}

func NewGetProject(projectKey string) (GetProject, error) {
	command := GetProject{
		ProjectKey: projectKey,
	}
	return command, command.validate()
}

func (c GetProject) validate() error {
	var result string
	if len(c.ProjectKey) == 0 {
		result += "Key must not be empty"
	}
	if len(result) != 0 {
		return errors.New(result)
	} else {
		return nil
	}
}

// Command for adding a comment to issue or project
type AddComment struct {
	ParentIssueKey string `json:"parrent_issue_key"`
	Content        string `json:"content"`
}

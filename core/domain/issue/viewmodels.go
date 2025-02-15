package issue

type CreateIssueCommand struct {
	IssueType  string `json:"issueType"`
	Name       string `json:"name"`
	ProjectKey string `json:"projectKey"`
}

package issue

import "time"

const (
	DDMMYYYYhhmmss = "2006-01-02 15:04:05"
)

// Read-only view of issues on issue list page: Tasks, Bugs, etc
type IssueListElement struct {
	ItemType   string `json:"item_type"`
	Name       string `json:"name"`
	ItemKey    string `json:"item_key"`
	State      string `json:"state"`
	AssignedTo string `json:"assigned_to"`
	CreatedBy  string `json:"created_by"`
	Created    string `json:"created"`
	Updated    string `json:updated"`
	ParentKey  string `json:"parent_key"`
}

func NewIssueListElement(id int, key, name, issueType, state, assagnee, creator string, created, updated time.Time, parentId int, parentKey string) IssueListElement {
	issue := IssueListElement{
		ItemType:   issueType,
		Name:       name,
		ItemKey:    key,
		State:      state,
		AssignedTo: assagnee,
		CreatedBy:  creator,
		Created:    created.Format(DDMMYYYYhhmmss),
		Updated:    updated.Format(DDMMYYYYhhmmss),
		ParentKey:  parentKey,
	}
	return issue
}

package internal


type Notification struct {
	RepositoryUri string `bson:"repositoryUri"`
	Filters       NotificationFilters
}

type NotificationFilters struct {
	Author		string
	Assignee	string
	Label 		string
	Title 		string
}

type GithubIssue struct {
	Title string `json:"title"`
}

package http


type Filters struct {
	Author		string `json:"author"`
	Assignee	string `json:"assignee"`
	Label 		string `json:"label"`
	Title 		string `json:"title"`
}

type NotificationCreationRequest struct {
	RepositoryUrl string 	`json:"repositoryUrl"`
	Email         string 	`json:"email"`
	Filters       Filters	`json:"filters"`
}

package http


type NotificationCreationRequest struct {
	RepositoryUrl string `json:"repositoryUrl"`
	Label         string `json:"label"`
	Email         string `json:"email"`
}

package app

type WorkspaceInfo struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	ImageUrl *string `json:"imageUrl"`
}

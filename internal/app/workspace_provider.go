package app

type WorkspaceProvider interface {
	Get(id string) (*WorkspaceInfo, error)
}

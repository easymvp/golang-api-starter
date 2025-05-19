package app

type UserProvider interface {
	Get(id string) (*UserInfo, error)
}

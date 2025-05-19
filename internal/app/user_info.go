package app

const IdentityKey = "userId"

type UserInfo struct {
	ID          string `json:"id"`
	WorkspaceID string `json:"workspaceId"`
	Username    string `json:"username"`
	FullName    string `json:"fullName"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	Status      string `json:"status"`
}

func (u UserInfo) IsAdmin() bool {
	return u.Role == "ADMIN"
}

func (u UserInfo) IsSuperAdmin() bool {
	return u.Role == "SUPER_ADMIN"
}

func (u UserInfo) IsActive() bool {
	return u.Role == "ACTIVE"
}

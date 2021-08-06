package models

type VehUserToGroupRequest struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	GroupID   string `json:"group_id"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"update_at"`
	DeleteAt  string `json:"delete_at"`
}

type VehUserToGroupUpdate struct {
	UserID    string `json:"user_id,omitempty"`
	GroupID   string `json:"group_id,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdateAt  string `json:"update_at,omitempty"`
	DeleteAt  string `json:"delete_at,omitempty"`
}

type VehUserToGroupResult struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	GroupID   string `json:"group_id"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"update_at"`
	DeleteAt  string `json:"delete_at"`
}

package models

type VehGroupRequest struct {
	ID             string `json:"id"`
	Name           string `json:"name" valid:"required"`
	OrganizationID string `json:"organization_id"`
	Description    string `json:"description"`
	CreateAt       string `json:"create_at"`
	UpdateAt       string `json:"update_at"`
	DeleteAt       string `json:"delete_at"`
}

type VehGroupUpdate struct {
	Name           string `json:"name,omitempty" valid:"required"`
	OrganizationID string `json:"organization_id,omitempty"`
	Description    string `json:"description,omitempty"`
	CreateAt       string `json:"create_at,omitempty"`
	UpdateAt       string `json:"update_at,omitempty"`
	DeleteAt       string `json:"delete_at,omitempty"`
}

type VehGroupResult struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	OrganizationID string `json:"organization_id"`
	Description    string `json:"description"`
	CreateAt       string `json:"create_at"`
	UpdateAt       string `json:"update_at"`
	DeleteAt       string `json:"delete_at"`
}

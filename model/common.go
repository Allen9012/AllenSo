package model

// ListParams 获取用户列表所需要绑定的参数
type ListParams struct {
	Page int
	Size int
}

type SearchDTO struct {
	Size int    `json:"size"`
	Page int    `json:"page"`
	Text string `json:"text"`
	Type string `json:"type"`
}

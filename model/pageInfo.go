package model

type PageInfo struct {
	PageSize  int            `json:"pageSize"`
	PageNum   int            `json:"pageNum"`
	OrderBy   string         `json:"orderBy"`
	LastPage  bool           `json:"lastPage"`
	Condition map[string]any `json:"condition"`
}

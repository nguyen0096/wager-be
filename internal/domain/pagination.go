package domain

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

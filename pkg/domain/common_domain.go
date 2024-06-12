package domain

type PagingPointer struct {
	Limit  int     `json:"limit"`
	LastID *string `json:"last_id"`
}

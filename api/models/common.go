package models

type GetListRequest struct {
	Page   int
	Limit  int
	Search string
}

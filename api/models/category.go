package models

type Category struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ParentID  string `json:"parent_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"-"`
}

type CreateCategory struct {
	Name     string `json:"name"`
	ParentID string `json:"parent_id"`
}

type UpdateCategory struct {
	ID        string `json:"-"`
	Name      string `json:"name"`
	ParentID  string `json:"parent_id"`
	UpdatedAt string `json:"-"`
}

type CategoryResponse struct {
	Categories []Category
	Count      int
}

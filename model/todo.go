package model

type TODO struct {
	Id          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
	IsActive    bool   `json:"is_active"`
}

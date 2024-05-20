package response

type PaginationReponse[T any] struct {
	Items T   `json:"items"`
	Total int `json:"total"`
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

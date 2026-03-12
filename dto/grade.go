package dto

type Grade struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Type  string  `json:"type"`
	Score float32 `json:"score"`
}

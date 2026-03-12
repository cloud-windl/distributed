package dto

type Student struct {
	ID        int     `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Grades    []Grade `json:"grades,omitempty"`
}

type StudentListResponse struct {
	List     []Student `json:"list"`
	Total    int64     `json:"total"`
	Page     int       `json:"page"`
	PageSize int       `json:"page_size"`
}

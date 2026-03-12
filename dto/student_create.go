package dto

type CreateStudentRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UpdateStudentRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

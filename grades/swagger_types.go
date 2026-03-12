package grades

import "distributed/dto"

type StudentListData struct {
	List     []dto.Student `json:"list"`
	Total    int64         `json:"total"`
	Page     int           `json:"page"`
	PageSize int           `json:"page_size"`
}

type StudentListResponseDoc struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data StudentListData `json:"data"`
}

type StudentResponseDoc struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data dto.Student `json:"data"`
}

type GradeListResponseDoc struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data []dto.Grade `json:"data"`
}

type GradeResponseDoc struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data dto.Grade `json:"data"`
}

type MessageData struct {
	Message string `json:"message"`
}

type MessageResponseDoc struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data MessageData `json:"data"`
}

type ErrorResponseDoc struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CreateStudentRequestDoc struct {
	FirstName string `json:"first_name" example:"Tom"`
	LastName  string `json:"last_name" example:"Jerry"`
}

type UpdateStudentRequestDoc struct {
	FirstName string `json:"first_name" example:"Alice"`
	LastName  string `json:"last_name" example:"Smith"`
}

type CreateGradeRequestDoc struct {
	Title string  `json:"title" example:"Math Final"`
	Type  string  `json:"type" example:"Exam"`
	Score float32 `json:"score" example:"95.5"`
}

type UpdateGradeRequestDoc struct {
	Title string  `json:"title" example:"Math Final"`
	Type  string  `json:"type" example:"Exam"`
	Score float32 `json:"score" example:"98"`
}

type CreateStudentRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UpdateStudentRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

package portal

import "distributed/dto"

type StudentsPageData struct {
	List     []dto.Student
	Total    int64
	Page     int
	PageSize int
	Keyword  string

	HasPrev  bool
	HasNext  bool
	PrevPage int
	NextPage int

	IsLoggedIn bool
}

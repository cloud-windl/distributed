package grades

type StudentRepo interface {
	GetAll() Students
	GetByID(id int) (*Student, error)
	AddGrade(id int, grade Grade) error
}

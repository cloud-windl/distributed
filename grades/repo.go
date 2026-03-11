package grades

type StudentRepo interface {
	GetAll(page, pageSize int, keyword string) (Students, int64, error)
	GetByID(id int) (*Student, error)
	Create(student *Student) (*Student, error)
	Update(id int, student *Student) (*Student, error)
	Delete(id int) error

	GetGradesByStudentID(studentID int) ([]Grade, error)
	AddGrade(id int, grade Grade) error

	GetGradeByID(id int) (*Grade, error)
	UpdateGrade(id int, grade Grade) (*Grade, error)
	DeleteGrade(id int) error
}

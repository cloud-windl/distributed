package grades

type StudentService interface {
	GetAllStudents() Students
	GetStudentByID(id int) (*Student, error)
	AddGrade(id int, grade Grade) error
}
type Service struct {
	repo StudentRepo
}

func NewService() StudentService {
	return &Service{}
}

func (s *Service) GetAllStudents() Students {
	return s.repo.GetAll()
}

func (s *Service) GetStudentByID(id int) (*Student, error) {
	return s.repo.GetByID(id)
}

func (s *Service) AddGrade(id int, grade Grade) error {
	return s.repo.AddGrade(id, grade)
}

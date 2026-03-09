package grades

type StudentService interface {
	GetAllStudents() Students
	GetStudentByID(id int) (*Student, error)
	AddGrade(id int, grade Grade) error
}
type Service struct{}

func NewService() StudentService {
	return &Service{}
}

func (s *Service) GetAllStudents() Students {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()

	return students
}

func (s *Service) GetStudentByID(id int) (*Student, error) {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()
	return students.GetByID(id)
}

func (s *Service) AddGrade(id int, grade Grade) error {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()

	student, err := students.GetByID(id)
	if err != nil {
		return err
	}

	student.Grades = append(student.Grades, grade)
	return nil
}

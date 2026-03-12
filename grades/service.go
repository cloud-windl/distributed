package grades

type StudentService interface {
	GetAllStudents(page, pageSize int, keyword string) (Students, int64, error)
	GetStudentByID(id int) (*Student, error)
	CreateStudent(student Student) (*Student, error)
	UpdateStudent(id int, student Student) (*Student, error)
	DeleteStudent(id int) error

	GetGradesByStudentID(studentID int) ([]Grade, error)
	AddGrade(studentID int, grade Grade) error
	GetGradeByID(id int) (*Grade, error)
	UpdateGrade(id int, grade Grade) (*Grade, error)
	DeleteGrade(id int) error
}

type Service struct {
	repo StudentRepo
}

func NewService(repo StudentRepo) StudentService {
	return &Service{repo: repo}
}

func (s *Service) GetAllStudents(page, pageSize int, keyword string) (Students, int64, error) {
	return s.repo.GetAll(page, pageSize, keyword)
}

func (s *Service) GetStudentByID(id int) (*Student, error) {
	return s.repo.GetByID(id)
}

func (s *Service) CreateStudent(student Student) (*Student, error) {
	return s.repo.Create(student)
}

func (s *Service) UpdateStudent(id int, student Student) (*Student, error) {
	return s.repo.Update(id, student)
}

func (s *Service) DeleteStudent(id int) error {
	return s.repo.Delete(id)
}

func (s *Service) GetGradesByStudentID(studentID int) ([]Grade, error) {
	return s.repo.GetGradesByStudentID(studentID)
}

func (s *Service) AddGrade(studentID int, grade Grade) error {
	return s.repo.AddGrade(studentID, grade)
}

func (s *Service) GetGradeByID(id int) (*Grade, error) {
	return s.repo.GetGradeByID(id)
}

func (s *Service) UpdateGrade(id int, grade Grade) (*Grade, error) {
	return s.repo.UpdateGrade(id, grade)
}

func (s *Service) DeleteGrade(id int) error {
	return s.repo.DeleteGrade(id)
}

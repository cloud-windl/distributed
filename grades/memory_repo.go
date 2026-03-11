package grades

type MemoryStudentRepo struct{}

func NewMemoryStudentRepo() *MemoryStudentRepo {
	return &MemoryStudentRepo{}
}

func (r *MemoryStudentRepo) GetAll() Students {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()
	return students
}

func (r *MemoryStudentRepo) GetByID(id int) (*Student, error) {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()
	return students.GetByID(id)
}
func (r *MemoryStudentRepo) AddGrade(id int, grade Grade) error {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()

	student, err := students.GetByID(id)
	if err != nil {
		return err
	}

	student.Grades = append(student.Grades, grade)
	return nil
}

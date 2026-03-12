package grades

import (
	"fmt"

	"gorm.io/gorm"
)

type MySQLStudentRepo struct {
	db *gorm.DB
}

func NewMySQLStudentRepo(db *gorm.DB) StudentRepo {
	return &MySQLStudentRepo{db: db}
}

func (r *MySQLStudentRepo) GetAll(page, pageSize int, keyword string) (Students, int64, error) {
	var models []StudentModel
	var total int64

	query := r.db.Model(&StudentModel{})

	if keyword != "" {
		like := "%" + keyword + "%"
		query = query.Where("first_name LIKE ? OR last_name LIKE ?", like, like)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	if err := query.Preload("Grades").
		Offset(offset).
		Limit(pageSize).
		Find(&models).Error; err != nil {
		return nil, 0, err
	}

	res := make(Students, 0, len(models))
	for _, m := range models {
		s := Student{
			ID:        int(m.ID),
			FirstName: m.FirstName,
			LastName:  m.LastName,
			Grades:    make([]Grade, 0, len(m.Grades)),
		}
		for _, g := range m.Grades {
			s.Grades = append(s.Grades, Grade{
				ID:    int(g.ID),
				Title: g.Title,
				Type:  GradeType(g.Type),
				Score: g.Score,
			})
		}
		res = append(res, s)
	}

	return res, total, nil
}

func (r *MySQLStudentRepo) GetByID(id int) (*Student, error) {
	var model StudentModel
	if err := r.db.Preload("Grades").First(&model, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("student with ID %d not found", id)
		}
		return nil, err
	}

	res := &Student{
		ID:        int(model.ID),
		FirstName: model.FirstName,
		LastName:  model.LastName,
		Grades:    make([]Grade, 0, len(model.Grades)),
	}
	for _, g := range model.Grades {
		res.Grades = append(res.Grades, Grade{
			ID:    int(g.ID),
			Title: g.Title,
			Type:  GradeType(g.Type),
			Score: g.Score,
		})
	}
	return res, nil
}

func (r *MySQLStudentRepo) Create(student *Student) (*Student, error) {
	model := StudentModel{
		FirstName: student.FirstName,
		LastName:  student.LastName,
	}
	if err := r.db.Create(&model).Error; err != nil {
		return nil, err
	}

	return &Student{
		ID:        int(model.ID),
		FirstName: model.FirstName,
		LastName:  model.LastName,
		Grades:    []Grade{},
	}, nil
}

func (r *MySQLStudentRepo) Update(id int, student *Student) (*Student, error) {
	var model StudentModel
	if err := r.db.First(&model, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("student with ID %d not found", id)
		}
		return nil, err
	}

	model.FirstName = student.FirstName
	model.LastName = student.LastName

	if err := r.db.Save(&model).Error; err != nil {
		return nil, err
	}

	return &Student{
		ID:        int(model.ID),
		FirstName: model.FirstName,
		LastName:  model.LastName,
	}, nil
}

func (r *MySQLStudentRepo) Delete(id int) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("student_id = ?", id).Delete(&GradeModel{}).Error; err != nil {
			return err
		}
		if err := tx.Delete(&StudentModel{}, id).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *MySQLStudentRepo) GetGradesByStudentID(studentID int) ([]Grade, error) {
	var models []GradeModel
	if err := r.db.Where("student_id = ?", studentID).Find(&models).Error; err != nil {
		return nil, err
	}

	res := make([]Grade, 0, len(models))
	for _, g := range models {
		res = append(res, Grade{
			ID:    int(g.ID),
			Title: g.Title,
			Type:  GradeType(g.Type),
			Score: g.Score,
		})
	}
	return res, nil
}

func (r *MySQLStudentRepo) AddGrade(studentID int, grade Grade) error {
	var student StudentModel
	if err := r.db.First(&student, studentID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("student with ID %d not found", studentID)
		}
		return err
	}

	g := GradeModel{
		StudentID: int64(studentID),
		Title:     grade.Title,
		Type:      string(grade.Type),
		Score:     grade.Score,
	}

	return r.db.Create(&g).Error
}

func (r *MySQLStudentRepo) GetGradeByID(id int) (*Grade, error) {
	var model GradeModel
	if err := r.db.First(&model, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("grade with ID %d not found", id)
		}
		return nil, err
	}

	return &Grade{
		ID:    int(model.ID),
		Title: model.Title,
		Type:  GradeType(model.Type),
		Score: model.Score,
	}, nil
}

func (r *MySQLStudentRepo) UpdateGrade(id int, grade Grade) (*Grade, error) {
	var model GradeModel
	if err := r.db.First(&model, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("grade with ID %d not found", id)
		}
		return nil, err
	}

	model.Title = grade.Title
	model.Type = string(grade.Type)
	model.Score = grade.Score

	if err := r.db.Save(&model).Error; err != nil {
		return nil, err
	}

	return &Grade{
		ID:    int(model.ID),
		Title: model.Title,
		Type:  GradeType(model.Type),
		Score: model.Score,
	}, nil
}

func (r *MySQLStudentRepo) DeleteGrade(id int) error {
	return r.db.Delete(&GradeModel{}, id).Error
}

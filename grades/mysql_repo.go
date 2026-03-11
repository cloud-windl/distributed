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

func (r *MySQLStudentRepo) GetAll() Students {
	var models []StudentModel
	r.db.Preload("Grades").Find(&models)

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
				Title: g.Title,
				Type:  GradeType(g.Type),
				Score: g.Score,
			})
		}
		res = append(res, s)
	}
	return res
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
			Title: g.Title,
			Type:  GradeType(g.Type),
			Score: g.Score,
		})
	}
	return res, nil
}

func (r *MySQLStudentRepo) AddGrade(id int, grade Grade) error {
	var student StudentModel
	if err := r.db.First(&student, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("student with ID %d not found", id)
		}
		return err
	}

	g := GradeModel{
		StudentID: int64(id),
		Title:     grade.Title,
		Type:      string(grade.Type),
		Score:     grade.Score,
	}

	return r.db.Create(&g).Error
}

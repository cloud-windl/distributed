package grades

import "distributed/dto"

func toGradeDTO(g *Grade) dto.Grade {
	return dto.Grade{
		ID:    g.ID,
		Title: g.Title,
		Type:  string(g.Type),
		Score: g.Score,
	}
}

func toGradeDTOList(grades []Grade) []dto.Grade {
	res := make([]dto.Grade, 0, len(grades))
	for i := range grades {
		res = append(res, dto.Grade{
			ID:    grades[i].ID,
			Title: grades[i].Title,
			Type:  string(grades[i].Type),
			Score: grades[i].Score,
		})
	}
	return res
}

func toStudentDTO(s *Student) dto.Student {
	res := dto.Student{
		ID:        s.ID,
		FirstName: s.FirstName,
		LastName:  s.LastName,
		Grades:    make([]dto.Grade, 0, len(s.Grades)),
	}

	for i := range s.Grades {
		res.Grades = append(res.Grades, dto.Grade{
			ID:    s.Grades[i].ID,
			Title: s.Grades[i].Title,
			Type:  string(s.Grades[i].Type),
			Score: s.Grades[i].Score,
		})
	}

	return res
}

func toStudentDTOList(students Students) []dto.Student {
	res := make([]dto.Student, 0, len(students))
	for i := range students {
		s := students[i]
		item := dto.Student{
			ID:        s.ID,
			FirstName: s.FirstName,
			LastName:  s.LastName,
			Grades:    make([]dto.Grade, 0, len(s.Grades)),
		}
		for j := range s.Grades {
			item.Grades = append(item.Grades, dto.Grade{
				ID:    s.Grades[j].ID,
				Title: s.Grades[j].Title,
				Type:  string(s.Grades[j].Type),
				Score: s.Grades[j].Score,
			})
		}
		res = append(res, item)
	}
	return res
}

package portal

import (
	"bytes"
	"distributed/grades"
	"distributed/registry"
	"encoding/json"
	"fmt"
	"net/http"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetStudents() (grades.Students, error) {
	var result grades.Students

	//serviceURL, err := registry.GetProvider(registry.GradingService)
	//if err != nil {
	//	return nil, err
	//}

	serviceURL := "http://localhost:6000"

	res, err := http.Get(serviceURL + "/students")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&result)
	return result, err
}

func (s *Service) GetStudentByID(id int) (grades.Student, error) {
	var result grades.Student

	serviceURL, err := registry.GetProvider(registry.GradingService)
	if err != nil {
		return result, err
	}

	res, err := http.Get(fmt.Sprintf("%v/students/%v", serviceURL, id))
	if err != nil {
		return result, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&result)
	return result, err
}

func (s *Service) AddGrade(id int, grade grades.Grade) error {
	data, err := json.Marshal(grade)
	if err != nil {
		return err
	}

	serviceURL, err := registry.GetProvider(registry.GradingService)
	if err != nil {
		return err
	}

	res, err := http.Post(fmt.Sprintf("%v/students/%v/grades", serviceURL, id), "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status code:%v", res.StatusCode)
	}

	return nil
}

package portal

import (
	"bytes"
	"distributed/grades"
	"distributed/pkg/response"
	"encoding/json"
	"fmt"
	"net/http"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

type studentsResponse struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data grades.Students `json:"data"`
}

type studentResponse struct {
	Code int            `json:"code"`
	Msg  string         `json:"msg"`
	Data grades.Student `json:"data"`
}

func (s *Service) GetStudents() (grades.Students, error) {
	var resp response.ClientResponse
	var result grades.Students

	serviceURL := "http://localhost:6000"

	res, err := http.Get(serviceURL + "/students")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf(resp.Msg)
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("empty response data")
	}

	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Service) GetStudentByID(id int) (grades.Student, error) {
	var resp response.ClientResponse
	var result grades.Student

	serviceURL := "http://localhost:6000"

	res, err := http.Get(fmt.Sprintf("%s/students/%d", serviceURL, id))
	if err != nil {
		return result, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return result, err
	}

	if resp.Code != 0 {
		return result, fmt.Errorf(resp.Msg)
	}

	if len(resp.Data) == 0 {
		return result, fmt.Errorf("empty response data")
	}

	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return result, err
	}

	return result, nil
}

func (s *Service) AddGrade(id int, g grades.Grade) error {
	var resp response.ClientResponse

	data, err := json.Marshal(g)
	if err != nil {
		return err
	}

	serviceURL := "http://localhost:6000"

	res, err := http.Post(
		fmt.Sprintf("%s/students/%d/grades", serviceURL, id),
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return err
	}

	if resp.Code != 0 {
		return fmt.Errorf(resp.Msg)
	}

	return nil
}

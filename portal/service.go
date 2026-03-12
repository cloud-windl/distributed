package portal

import (
	"bytes"
	"distributed/dto"
	"distributed/pkg/response"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetStudents(page, pageSize int, keyword string) (dto.StudentListResponse, error) {
	var resp response.ClientResponse
	var result dto.StudentListResponse

	serviceURL := "http://localhost:6001"

	q := url.Values{}
	q.Set("page", strconv.Itoa(page))
	q.Set("page_size", strconv.Itoa(pageSize))
	q.Set("keyword", keyword)

	reqURL := fmt.Sprintf("%s/students?%s", serviceURL, q.Encode())

	res, err := http.Get(reqURL)
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

func (s *Service) GetStudentByID(id int) (dto.Student, error) {
	var resp response.ClientResponse
	var result dto.Student

	serviceURL := "http://localhost:6001"

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

func (s *Service) AddGrade(id int, g dto.Grade) error {
	var resp response.ClientResponse

	data, err := json.Marshal(g)
	if err != nil {
		return err
	}

	serviceURL := "http://localhost:6001"

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

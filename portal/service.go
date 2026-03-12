package portal

import (
	"bytes"
	"distributed/dto"
	"distributed/pkg/response"
	"distributed/registry"
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

func (s *Service) gradesServiceURL() (string, error) {
	return registry.GetProvider(registry.GradingService)
}

func (s *Service) Login(username, password string) (string, error) {
	serviceURL, err := s.gradesServiceURL()
	if err != nil {
		return "", err
	}

	body := map[string]string{
		"username": username,
		"password": password,
	}

	data, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	res, err := http.Post(
		serviceURL+"/login",
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	var resp response.ClientResponse
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return "", err
	}

	if resp.Code != 0 {
		return "", fmt.Errorf(resp.Msg)
	}

	var result struct {
		Token string `json:"token"`
	}

	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return "", err
	}

	return result.Token, nil
}

func (s *Service) GetStudents(page, pageSize int, keyword, token string) (dto.StudentListResponse, error) {
	var resp response.ClientResponse
	var result dto.StudentListResponse

	serviceURL, err := s.gradesServiceURL()
	if err != nil {
		return result, err
	}

	q := url.Values{}
	q.Set("page", strconv.Itoa(page))
	q.Set("page_size", strconv.Itoa(pageSize))
	q.Set("keyword", keyword)

	reqURL := fmt.Sprintf("%s/students?%s", serviceURL, q.Encode())

	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return result, err
	}
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	client := &http.Client{}
	res, err := client.Do(req)
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

func (s *Service) GetStudentByID(id int, token string) (dto.Student, error) {
	var resp response.ClientResponse
	var result dto.Student

	serviceURL, err := s.gradesServiceURL()
	if err != nil {
		return result, err
	}

	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/students/%d", serviceURL, id),
		nil,
	)
	if err != nil {
		return result, err
	}
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	client := &http.Client{}
	res, err := client.Do(req)
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

func (s *Service) CreateStudent(firstName, lastName, token string) error {
	var resp response.ClientResponse

	serviceURL, err := s.gradesServiceURL()
	if err != nil {
		return err
	}

	body := map[string]string{
		"first_name": firstName,
		"last_name":  lastName,
	}

	data, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/students", serviceURL),
		bytes.NewBuffer(data),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	res, err := client.Do(req)
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

func (s *Service) UpdateStudent(id int, firstName, lastName, token string) error {
	var resp response.ClientResponse

	serviceURL, err := s.gradesServiceURL()
	if err != nil {
		return err
	}

	body := map[string]string{
		"first_name": firstName,
		"last_name":  lastName,
	}

	data, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPut,
		fmt.Sprintf("%s/students/%d", serviceURL, id),
		bytes.NewBuffer(data),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	res, err := client.Do(req)
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

func (s *Service) DeleteStudent(id int, token string) error {
	var resp response.ClientResponse

	serviceURL, err := s.gradesServiceURL()
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("%s/students/%d", serviceURL, id),
		nil,
	)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	res, err := client.Do(req)
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

func (s *Service) AddGrade(id int, g dto.Grade, token string) error {
	var resp response.ClientResponse

	serviceURL, err := s.gradesServiceURL()
	if err != nil {
		return err
	}

	data, err := json.Marshal(g)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/students/%d/grades", serviceURL, id),
		bytes.NewBuffer(data),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	res, err := client.Do(req)
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

func (s *Service) UpdateGrade(id int, g dto.Grade, token string) error {
	var resp response.ClientResponse

	serviceURL, err := s.gradesServiceURL()
	if err != nil {
		return err
	}

	data, err := json.Marshal(g)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPut,
		fmt.Sprintf("%s/grades/%d", serviceURL, id),
		bytes.NewBuffer(data),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	res, err := client.Do(req)
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

func (s *Service) DeleteGrade(id int, token string) error {
	var resp response.ClientResponse

	serviceURL, err := s.gradesServiceURL()
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("%s/grades/%d", serviceURL, id),
		nil,
	)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	res, err := client.Do(req)
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

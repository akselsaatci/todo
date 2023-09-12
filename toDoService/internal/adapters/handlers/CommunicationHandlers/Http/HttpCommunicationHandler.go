package Http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"toDoService/internal/core/dtos"
)

type HttpCommunicationHandler struct {
}

func NewHttpCommunicationHandler() *HttpCommunicationHandler {
	return &HttpCommunicationHandler{}
}

func (s *HttpCommunicationHandler) AuthorizeUserByToken(token string) (*dtos.UserDTO, error) {
	// implemented very bad im gonna change these
	requestURL := "http://localhost:8080"

	requestData := map[string]string{
		"token": token,
	}

	requestDataBytes, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}

	requestBody := bytes.NewBuffer(requestDataBytes)

	res, err := http.Post(requestURL, "application/json", requestBody)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		// :P
		return nil, errors.New(fmt.Sprint(res.StatusCode))
	}

	var userDTO dtos.UserDTO
	if err := json.NewDecoder(res.Body).Decode(&userDTO); err != nil {
		return nil, err
	}

	return &userDTO, nil
}

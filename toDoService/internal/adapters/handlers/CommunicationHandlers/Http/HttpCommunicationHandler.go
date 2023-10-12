package Http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"toDoService/internal/core/customErrors"
	"toDoService/internal/core/dtos"
)

type HttpCommunicationHandler struct {
	AUTH_URL string
}

func NewHttpCommunicationHandler(auth_URL string) *HttpCommunicationHandler {
	return &HttpCommunicationHandler{
		AUTH_URL: auth_URL,
	}
}

func (s *HttpCommunicationHandler) AuthorizeUserByToken(token string) (*dtos.UserDTO, error) {
	// implemented very poorly im gonna change these
	baseRequestPath := s.AUTH_URL
	requestURL, err := url.JoinPath(baseRequestPath, "/validate")
	if err != nil {
		// TODO handle this error
		log.Fatalf("%s", err)
	}

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

	if res.StatusCode == 401 {
		err = &customErrors.UnauthorizedError{}
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		//TODO what is the better way to handle this error ? im probablily not going to use http to commincate between services so it can stay fn
		err = errors.New(fmt.Sprint(res.StatusCode))
		return nil, err
	}

	var userDTO dtos.UserDTO
	if err := json.NewDecoder(res.Body).Decode(&userDTO); err != nil {
		return nil, err
	}

	return &userDTO, nil
}

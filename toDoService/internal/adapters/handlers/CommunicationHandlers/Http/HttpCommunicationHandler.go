package Http

import "toDoService/internal/core/dtos"

type HttpCommunicationHandler struct {
}

func NewHttpCommunicationHandler() *HttpCommunicationHandler {
	return &HttpCommunicationHandler{}
}

func (s *HttpCommunicationHandler) AuthorizeUserByToken(token string) (*dtos.UserDTO, *error) {

}

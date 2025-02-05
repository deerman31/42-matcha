package otherusers

type OtherUsersHandler struct {
	service *OtherUsersService
}

func NewOtherUsersHandler(service *OtherUsersService) *OtherUsersHandler {
	return &OtherUsersHandler{service: service}
}

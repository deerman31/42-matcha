package myprofile

type MyProfileHandler struct {
	service *MyProfileService
}

func  NewMyProfileHandler(service *MyProfileService) *MyProfileHandler {
	return &MyProfileHandler{service: service}
}

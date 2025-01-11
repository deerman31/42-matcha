package dev

type FiveThousandRegisterHandler struct {
	service *FiveThousandRegisterService
}

func NewFiveThousandRegisterHandler(service *FiveThousandRegisterService) *FiveThousandRegisterHandler {
	return &FiveThousandRegisterHandler{service: service}
}

package fakeaccount

type FakeAccountHandler struct {
	service *FakeAccountService
}

func NewFakeAccountHandler(service *FakeAccountService) *FakeAccountHandler {
	return &FakeAccountHandler{service: service}
}

package research

type ResearchHandler struct {
	service *ResearchService
}

func NewResearchHandler(service *ResearchService) *ResearchHandler {
	return &ResearchHandler{service: service}
}

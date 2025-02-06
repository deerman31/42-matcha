package tags

type SetTagResponse struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

type SearchTagResponse struct {
	Tags  []string `json:"tags,omitempty"`
	Error string   `json:"error,omitempty"`
}

type GetUserTagResponse struct {
	Tags  []string `json:"tags,omitempty"`
	Error string   `json:"error,omitempty"`
}

type DeleteTagResponse struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}
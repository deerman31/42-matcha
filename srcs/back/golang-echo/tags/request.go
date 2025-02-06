package tags

type SetTagRequest struct {
	Tag string `json:"tag" validate:"required,tag"`
}

type SearchTagRequest struct {
	TagName string `json:"tagname" validate:"required"`
}

type DeleteTagRequest struct {
	Tag string `json:"tag" validate:"required,tag"`
}
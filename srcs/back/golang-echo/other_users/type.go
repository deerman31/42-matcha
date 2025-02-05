package otherusers

type OtherGetImageRequest struct {
	ImagePath string `json:"image_path" validate:"required"`
}

type OtherGetImageResponse struct {
	Image string `json:"image,omitempty"`
	Error string `json:"error,omitempty"`
}

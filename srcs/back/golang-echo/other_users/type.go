package otherusers

type OtherGetImageRequest struct {
	ImagePath string `json:"image_path" validate:"required"`
}

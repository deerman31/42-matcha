package otherusers

import (
	"golang-echo/app/utils"
	"golang-echo/app/utils/query"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (o *OtherUsersHandler) GetOtherAllImage(c echo.Context) error {
	username := c.Param("name")
	if err := c.Validate(struct {
		Username string `json:"username" validate:"required,username"`
	}{Username: username}); err != nil {
		return c.JSON(http.StatusBadRequest, GetOtherAllImageResponse{Error: err.Error()})
	}
	allImage, err := o.service.GetOtherAllImage(username)
	if err != nil {
		switch err {
		case ErrTransactionFailed:
			return c.JSON(http.StatusInternalServerError, GetOtherAllImageResponse{Error: err.Error()})
		default:
			return c.JSON(http.StatusInternalServerError, GetOtherAllImageResponse{Error: err.Error()})
		}
	}
	return c.JSON(http.StatusOK, GetOtherAllImageResponse{AllImage: allImage})
}

func (o *OtherUsersService) GetOtherAllImage(userName string) ([]string, error) {
	tx, err := o.db.Begin()
	if err != nil {
		return nil, ErrTransactionFailed
	}
	defer tx.Rollback()
	otherID, err := query.GetUserIDByUsername(tx, userName)
	if err != nil {
		return nil, err
	}
	allImagePaths, err := query.GetAllImagePath(tx, otherID)
	if err != nil {
		return nil, err
	}
	retImages, err := utils.SetAllImageURI(allImagePaths)
	if err != nil {
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		return nil, ErrTransactionFailed
	}
	return retImages, nil

}

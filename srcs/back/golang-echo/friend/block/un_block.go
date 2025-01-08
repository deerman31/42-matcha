package block

import (
	"database/sql"
	"golang-echo/jwt_token"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (b *BlockHandler) UnBlock(c echo.Context) error {
	claims, ok := c.Get("user").(*jwt_token.Claims)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "user claims not found")
	}
	req := new(BlockRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, BlockResponse{Error: "Invalid request body"})
	}
	// validationをここで行う
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, BlockResponse{Error: err.Error()})
	}

	err := b.blockService.UnBlock(claims.UserID, req.Username)
	if err != nil {
		switch err {
		case ErrUserNotFound:
			return c.JSON(http.StatusNotFound, BlockResponse{Error: "User not found"})
		case ErrSelfAction:
			return c.JSON(http.StatusBadRequest, BlockResponse{Error: "Cannot unblock your own profile"})
		case ErrBlockNotFound:
			return c.JSON(http.StatusNotFound, BlockResponse{Error: "Block not found"})
		default:
			return c.JSON(http.StatusInternalServerError, BlockResponse{Error: "Internal server error"})
		}
	}

	return c.JSON(http.StatusNoContent, BlockResponse{Message: "Successfully unblock the user"})
}

func (b *BlockService) UnBlock(myID int, userName string) error {
	tx, err := b.db.Begin()
	if err != nil {
		return ErrTransactionFailed
	}
	defer tx.Rollback() // エラーが発生した場合はロールバック

	otherID, err := getUserID(tx, userName)
	if err != nil {
		return err
	}
	if myID == otherID {
		return ErrSelfAction
	}

	if err := deleteBlock(tx, myID, otherID); err != nil {
		return err
	}
	return tx.Commit()
}

func deleteBlock(tx *sql.Tx, myID, otherID int) error {
	const queryDeleteBlock = `
		DELETE FROM user_blocks
		WHERE user_id1 = $1 AND user_id2 = $2
	`
	result, err := tx.Exec(queryDeleteBlock, myID, otherID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrBlockNotFound
	}
	return nil
}

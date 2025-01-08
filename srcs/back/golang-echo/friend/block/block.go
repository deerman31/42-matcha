package block

import (
	"database/sql"
	"golang-echo/jwt_token"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func (b *BlockHandler) Block(c echo.Context) error {
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

	err := b.blockService.Block(claims.UserID, req.Username)
	if err != nil {
		switch err {
		case ErrUserNotFound:
			return c.JSON(http.StatusNotFound, BlockResponse{Error: "User not found"})
		case ErrSelfAction:
			return c.JSON(http.StatusBadRequest, BlockResponse{Error: "Cannot block your own profile"})
		case ErrAlreadyBlock:
			return c.JSON(http.StatusConflict, BlockResponse{Error: "Already block with this user"})
		default:
			return c.JSON(http.StatusInternalServerError, BlockResponse{Error: "Internal server error"})
		}
	}
	return c.JSON(http.StatusCreated, BlockResponse{Message: "Successfully blocked the user"})
}

func (b *BlockService) Block(myID int, userName string) error {
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
	if err := createBlock(tx, myID, otherID); err != nil {
		return err
	}
	return tx.Commit()
}

func createBlock(tx *sql.Tx, blockerID, blockedID int) error {
	const queryInsertBlock = `
	INSERT INTO user_blocks (blocker_id, blocked_id)
	VALUES ($1, $2)
	`
	_, err := tx.Exec(queryInsertBlock, blockerID, blockedID)
	if pqErr, ok := err.(*pq.Error); ok {
		if pqErr.Code == "23505" { // unique_violation のエラーコード
			return ErrAlreadyBlock
		}
	}
	return err
}

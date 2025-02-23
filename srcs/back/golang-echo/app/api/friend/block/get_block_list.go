package block

import (
	"golang-echo/app/utils/jwt_token"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (b *BlockHandler) GetBlockList(c echo.Context) error {
	claims, ok := c.Get("user").(*jwt_token.Claims)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "user claims not found")
	}
	userNames, err := b.blockService.GetBlockList(claims.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, BlockResponse{Error: "Internal server error"})
	}
	return c.JSON(http.StatusOK, BlockListResponse{Users: userNames})

}

// 自分がblockしたuserのlist
func (b *BlockService) GetBlockList(myID int) ([]string, error) {
	const query = `
SELECT u.username 
FROM users u
INNER JOIN user_blocks ub ON u.id = ub.blocked_id 
WHERE ub.blocker_id = $1
ORDER BY ub.created_at DESC;`
	rows, err := b.db.Query(query, myID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usernames []string
	for rows.Next() {
		var username string
		if err := rows.Scan(&username); err != nil {
			return nil, err
		}
		usernames = append(usernames, username)
	}
	return usernames, rows.Err()
}

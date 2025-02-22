package browse

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func BrowseRoutes(protected *echo.Group, db *sql.DB) {
	browse := protected.Group("/browse")

	handler := NewBrowseHandler(NewBrowseService(db))

	browse.POST("", handler.GetBrowseUser)
}

package research

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func ResearchRoutes(protected *echo.Group, db *sql.DB) {
	browse := protected.Group("/research")

	handler := NewResearchHandler(NewResearchService(db))

	browse.GET("", handler.GetResearchUsers)
}

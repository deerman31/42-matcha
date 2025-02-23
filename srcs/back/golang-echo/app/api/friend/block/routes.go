package block

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func BlockRoutes(protected *echo.Group, db *sql.DB) {
	block := protected.Group("/block")

	blockHandler := NewBlockHandler(NewBlockService(db))
	block.POST("/block", blockHandler.Block)
	block.DELETE("/un-block", blockHandler.UnBlock)
	block.GET("/get-block-list", blockHandler.GetBlockList)
}

package block

type BlockHandler struct {
	blockService *BlockService
}

func NewBlockHandler(blockService *BlockService) *BlockHandler {
	return &BlockHandler{blockService: blockService}
}

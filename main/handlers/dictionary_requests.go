package handlers

type dictionaryItemRequest struct {
	Type  string `json:"type" binding:"required"`
	Label string `json:"label" binding:"required"`
}

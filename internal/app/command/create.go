package command

import (
	"context"

	"github.com/trevatk/gotmpl/internal/domain/entity"
)

// CreateHandler
type CreateHandler struct {
	repo entity.IRepository
}

// NewCreateHandler
func NewCreateHandler(repo entity.IRepository) *CreateHandler {
	return &CreateHandler{repo: repo}
}

// Handle
func (h *CreateHandler) Handle(ctx context.Context) error {
	return nil
}

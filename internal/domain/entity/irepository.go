package entity

import "context"

// IRepository
type IRepository interface {
	Create(ctx context.Context, e *Entity) error
}

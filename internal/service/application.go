package service

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/trevatk/common/database"

	"github.com/trevatk/gotmpl/internal/adapter"
	"github.com/trevatk/gotmpl/internal/app"
	"github.com/trevatk/gotmpl/internal/app/command"
)

// ProvideApplication
func ProvideApplication(lc fx.Lifecycle, log *zap.Logger) (*app.Application, error) {

	cfg := database.Config{}
	db, err := cfg.ProvideDatabase(log)
	if err != nil {
		return nil, err
	}

	repo := adapter.ProvideRepository(db)

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return nil
			},
		},
	)

	return &app.Application{
		Commands: &app.Commands{
			CreateHandler: command.NewCreateHandler(repo),
		},
	}, nil
}

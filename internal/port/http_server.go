package port

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/fx"

	"github.com/trevatk/common/server"

	"github.com/trevatk/gotmpl/internal/app"
	"github.com/trevatk/gotmpl/internal/config"
)

const (
	_entityRoute = "/entity"
)

// IHttpServer
type IHttpServer interface {
	createEntity(w http.ResponseWriter, r *http.Request)
}

// HttpServer
type HttpServer struct {
	app *app.Application
}

var _ IHttpServer = (*HttpServer)(nil)

// InvokeHttpServer
func InvokeHttpServer(lc fx.Lifecycle, cfg *config.Config, app *app.Application) error {

	s, err := server.ProvideServer(cfg.Server.Port())
	if err != nil {
		return err
	}

	hs := &HttpServer{app: app}

	srv := &http.Server{
		Handler: registerRouter(hs),
	}

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				return s.ServeHTTP(srv)
			},
			OnStop: func(ctx context.Context) error {
				return s.Shutdown(ctx, srv)
			},
		},
	)

	return nil
}

func (hs *HttpServer) createEntity(w http.ResponseWriter, r *http.Request) {}

func registerRouter(handler IHttpServer) *mux.Router {

	router := mux.NewRouter()

	v1 := router.PathPrefix("/api/v1").Subrouter()

	v1.Path(_entityRoute).HandlerFunc(handler.createEntity).Methods(http.MethodPost)

	return router
}

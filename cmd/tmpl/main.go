package main

import (
	"context"
	"log"
	"time"

	"go.uber.org/fx"

	"github.com/trevatk/common/logging"

	"github.com/trevatk/gotmpl/internal/config"
	"github.com/trevatk/gotmpl/internal/port"
	"github.com/trevatk/gotmpl/internal/service"
)

func main() {

	app := fx.New(
		fx.Provide(logging.ProvideLogger),
		fx.Provide(config.ProvideConfig),
		fx.Invoke(config.InvokeConfig),
		fx.Provide(service.ProvideApplication),
		fx.Invoke(port.InvokeHttpServer),
	)

	start, cancel := context.WithTimeout(context.TODO(), time.Second*15)
	defer cancel()

	if err := app.Start(start); err != nil {
		log.Fatal(err)
	}

	<-app.Done()

	stop, cancel := context.WithTimeout(context.TODO(), time.Second*15)
	defer cancel()

	if err := app.Stop(stop); err != nil {
		log.Fatal(err)
	}
}

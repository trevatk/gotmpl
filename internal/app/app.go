package app

import "github.com/trevatk/gotmpl/internal/app/command"

// Application class to hold all command/query handlers
type Application struct {
	Commands *Commands
	Queries  *Queries
}

// Commands class to hold all application command handlers
type Commands struct {
	CreateHandler *command.CreateHandler
}

// Queries class to hold all application query handlers
type Queries struct{}

package controller

import (
	"context"
)

type App struct {
	db *storage.DB
}

func NewApp(store *storage.DB) *App {
	a := &App{
		db: store,
	}
	return a
}

type APIServer interface {
	Start(db *storage.DB)
	Stop()
}

func (a *App) Serve(ctx context.Context, hs APIServer) {
	hs.Start(a.db)
	<-ctx.Done()
	hs.Stop()
}

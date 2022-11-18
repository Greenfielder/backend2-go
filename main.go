package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/Greenfielder/backend2-go/internal/api/handler"
	"github.com/Greenfielder/backend2-go/internal/api/server"
	"github.com/Greenfielder/backend2-go/internal/entity"
	"github.com/Greenfielder/backend2-go/internal/controller"
	"github.com/Greenfielder/backend2-go/internal/repository"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	connString := `postgresql://user:user@127.0.0.1:5432/pet_db`
	db, err := postgres.NewDB(ctx, connString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := storage.NewDB(db)
	a := starter.NewApp(repo)
	h := handler.NewRouter(repo)
	srv := server.NewServer(":8000", h)

	wg := &sync.WaitGroup{}
	go func() {
		wg.Add(1)
		a.Serve(ctx, srv)
		wg.Done()
	}()

	select {
	case err := <-srv.Err:
		log.Println(fmt.Errorf("server error: %w", err))
	case <-ctx.Done():
	}
	cancel()
	wg.Wait()
}

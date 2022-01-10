package main

import (
	"betspoiler_test/internal"
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
)

func main()  {

	router := chi.NewRouter()
	log.Println("router created")
	router.Use(middleware.Logger)
	log.Println("logger started")

	ctx, cancel := context.WithCancel(context.Background())
	internal.StopSignal(cancel)
	internal.Start(ctx, router)

	log.Println("main: done")
}
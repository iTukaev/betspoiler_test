package internal

import (
	"betspoiler_test/pkg/chatServise"
	"betspoiler_test/pkg/handlers/addMessage"
	"betspoiler_test/pkg/handlers/getAll"
	"context"
	"github.com/go-chi/chi"
	"log"
	"net"
	"net/http"
	"time"
)

func Start(ctx context.Context, r *chi.Mux) {

	chat := chatServise.NewChat()
	r.Get("/", getAll.NewHandler(chat))
	r.Post("/", addMessage.NewHandler(chat))

	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Println(err)
	}

	server := &http.Server{
		Handler: r,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Println("server started")
	go func() {
		log.Println(server.Serve(listener))
	}()

	<-ctx.Done()
	if err := server.Close(); err != nil {
		log.Println(err)
		return
	}
	log.Println("server stopped")
}


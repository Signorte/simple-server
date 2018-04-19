package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Signorte/serverTutorial/project-layout/api"
)

func main() {
	mux := http.NewServeMux()

	api.Register(mux)

	exit := make(chan os.Signal)

	signal.Notify(exit, os.Interrupt)

	server := http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	go func() { server.ListenAndServe() }()

	<-exit

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	server.Shutdown(ctx)
}

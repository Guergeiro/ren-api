package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/guergeiro/fator-conversao-gas-portugal/cmd/v1"
	"github.com/rs/cors"
)

func main() {
	port, ok := os.LookupEnv("PORT")
	if ok == false {
		port = "49152"
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	v1Router := http.NewServeMux()
	v1Router.Handle("/v1/", http.StripPrefix("/v1", v1.NewRoutes()))

	handler := cors.Default().Handler(v1Router)

	defer stop()
	server := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", port),
		Handler: handler,
	}

	/*
		Initializing the server in a goroutine so that
		it won't block the graceful shutdown handling below
	*/
	go func() {
		log.Println(fmt.Sprintf("Listening on http://%s", server.Addr))
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			// it is fine to use Fatal here because it is not main gorutine
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	/*
		Restore default behavior on the interrupt signal and notify user of
		shutdown.
	*/
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	/*
		The context is used to inform the server it has 5 seconds to finish
		the request it is currently handling
	*/
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

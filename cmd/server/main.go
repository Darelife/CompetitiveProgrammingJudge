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

	"github.com/darelife/competitiveprogrammingjudge/internal/transport"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/submit", transport.HandleSubmission)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// err := http.ListenAndServe(":8080", mux)
	// if err != nil {
	// 	fmt.Printf("Error starting server: %s\n", err)
	// }

	go func() {
		fmt.Println("Server is running on port 8080")
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen error: %s\n", err)
		}
	}()

	<-done
	fmt.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatalf("Server shutdown error: %s\n", err)
	}
	fmt.Println("Server gracefully stopped")
}

// func handleSubmission(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Submission received!")
// }

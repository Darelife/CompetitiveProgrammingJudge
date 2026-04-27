// ============================================================
// FILE: cmd/server/main.go
// ROLE: Entry point for the judge HTTP server.
//
// FLOW:
//   1. Register a single route:  POST /submit  →  transport.HandleSubmission
//   2. Start listening on :8080
//   3. Every incoming request is handled by the transport layer (see
//      internal/transport/handlers.go) which does the actual judging.
//
// PIPELINE POSITION:
//   [Client] ──HTTP POST /submit──▶ [Server :8080] ──▶ [Handler]
// ============================================================
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/darelife/competitiveprogrammingjudge/internal/transport"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/submit", transport.HandleSubmission)

	fmt.Println("Judge server running on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

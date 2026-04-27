// ============================================================
// FILE: internal/transport/handlers.go
// ROLE: HTTP layer. Parses every POST /submit request, delegates to judge,
//       and returns the HTTP response.
//
// STEP-BY-STEP FLOW inside HandleSubmission():
//
//   [1] Decode JSON body  ──▶  extract req.Code
//         └─ if missing/invalid ──▶ 400 Bad Request
//
//   [2] Delegate to judge package  ──▶  judge.Evaluate(req.Code)
//         └─ it handles compiling, diffing, executing.
//
//   [3] Send HTTP response  ──▶  200 OK + SubmissionResponse {verdict}
//         └─ if judge errors  ──▶ sends error code + error message
// ============================================================
package transport

import (
	"encoding/json"
	"net/http"

	"github.com/darelife/competitiveprogrammingjudge/internal/judge"
	"github.com/darelife/competitiveprogrammingjudge/internal/models"
)

// HandleSubmission parses the incoming request and delegates
// the compilation and execution to the judge package.
func HandleSubmission(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.SubmissionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Code == "" {
		sendJSONError(w, "Invalid request: 'code' field required", http.StatusBadRequest)
		return
	}

	// Delegate judging logic to external package for cleaner separation
	resp, judgeErr := judge.Evaluate(req.Code)
	if judgeErr != nil {
		sendJSONError(w, judgeErr.Message, judgeErr.Code)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func sendJSONError(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

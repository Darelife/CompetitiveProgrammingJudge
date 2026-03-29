package transport

import (
	"encoding/json"
	"net/http"

	"github.com/darelife/competitiveprogrammingjudge/internal/models"
)

func HandleSubmission(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.SubmissionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		sendJSONError(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if req.Code == "" || req.Language == "" {
		sendJSONError(w, "Code and language are required", http.StatusBadRequest)
		return
	}

	res := models.SubmissionResponse{
		Message:      "Submission received successfully",
		SubmissionID: "0",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(202)
	json.NewEncoder(w).Encode(res)
}

func sendJSONError(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

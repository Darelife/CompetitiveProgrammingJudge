// ============================================================
// FILE: internal/models/submission.go
// ROLE: Shared data types — the wire format between client and server.
//
// DATA FLOW:
//
//   CLIENT ──▶ JSON encode ──▶ HTTP body ──▶ SERVER decodes into:
//
//   SubmissionRequest
//     └─ Code  string   raw C++ source code to judge
//
//   SERVER judges, then JSON encodes and responds with:
//
//   SubmissionResponse
//     ├─ Verdict  "AC" | "WA"
//     ├─ Output   actual stdout produced by the submitted program
//     └─ Diff     (only on WA) human-readable line-by-line diff
// ============================================================
package models

type SubmissionRequest struct {
	Code string `json:"code"`
}

type SubmissionResponse struct {
	Verdict string `json:"verdict"`          // "AC" or "WA"
	Diff    string `json:"diff,omitempty"`   // only on WA
	Output  string `json:"output,omitempty"` // actual output
}

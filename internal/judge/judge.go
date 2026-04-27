// ============================================================
// FILE: internal/judge/judge.go
// ROLE: Core judging logic. Sandbox execution, compilation, diffing.
//
// STEP-BY-STEP FLOW inside Evaluate():
//
//	[1] Create a temp directory  (e.g. /tmp/judge-abc123/)
//	      ├─ solution.cpp   ← write req.Code here
//	      └─ solution       ← compiled binary will go here
//
//	[2] Compile  ──▶  g++ -o solution solution.cpp
//	      └─ if g++ fails ──▶ 422 Unprocessable Entity  +  compiler stderr
//
//	[3] Run  ──▶  ./solution  <  data/input.txt
//	      └─ if runtime error ──▶ 422 Unprocessable Entity
//
//	[4] Compare output
//	      actual   = stdout of ./solution   (trimmed)
//	      expected = data/expected_output.txt (trimmed)
//
//	[5] Verdict
//	      actual == expected  ──▶ { verdict: "AC" }
//	      actual != expected  ──▶ { verdict: "WA", diff: buildDiff(...) }
//
// HELPER: buildDiff(expected, actual)
//
//	Walks both strings line-by-line and reports every differing line.
//
// ============================================================
package judge

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/darelife/competitiveprogrammingjudge/internal/models"
)

// dataDir is resolved relative to the working directory of the server binary.
const dataDir = "data"

// Error wraps an HTTP status code and a message.
type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

// Evaluate takes the C++ source code, compiles it, runs it against the input,
// and compares the output with expected output.
func Evaluate(code string) (*models.SubmissionResponse, *Error) {
	// 1. Setup workspace
	tmpDir, err := os.MkdirTemp("", "judge-*")
	if err != nil {
		return nil, &Error{500, "Failed to create temp dir"}
	}
	defer os.RemoveAll(tmpDir)

	srcFile := filepath.Join(tmpDir, "solution.cpp")
	binFile := filepath.Join(tmpDir, "solution")

	// 2. Write source code
	if err := os.WriteFile(srcFile, []byte(code), 0644); err != nil {
		return nil, &Error{500, "Failed to write source file"}
	}

	// 3. Compile
	compileCmd := exec.Command("g++", "-o", binFile, srcFile)
	if compileOut, err := compileCmd.CombinedOutput(); err != nil {
		return nil, &Error{422, "Compilation error:\n" + string(compileOut)}
	}

	// 4. Run
	inputBytes, err := os.ReadFile(filepath.Join(dataDir, "input.txt"))
	if err != nil {
		return nil, &Error{500, "Server error: could not read input"}
	}

	runCmd := exec.Command(binFile)
	runCmd.Stdin = strings.NewReader(string(inputBytes))
	actualOut, err := runCmd.Output()
	if err != nil {
		return nil, &Error{422, "Runtime error: " + err.Error()}
	}

	// 5. Compare & generate verdict
	expectedBytes, err := os.ReadFile(filepath.Join(dataDir, "expected_output.txt"))
	if err != nil {
		return nil, &Error{500, "Server error: could not read expected output"}
	}

	actual := strings.TrimSpace(string(actualOut))
	expected := strings.TrimSpace(string(expectedBytes))

	resp := &models.SubmissionResponse{Output: actual}
	if actual == expected {
		resp.Verdict = "AC"
	} else {
		resp.Verdict = "WA"
		resp.Diff = buildDiff(expected, actual)
	}

	return resp, nil
}

// buildDiff produces a simple line-by-line diff string.
func buildDiff(expected, actual string) string {
	expLines := strings.Split(expected, "\n")
	actLines := strings.Split(actual, "\n")

	var sb strings.Builder
	maxLen := len(expLines)
	if len(actLines) > maxLen {
		maxLen = len(actLines)
	}
	for i := 0; i < maxLen; i++ {
		exp, act := "", ""
		if i < len(expLines) {
			exp = expLines[i]
		}
		if i < len(actLines) {
			act = actLines[i]
		}
		if exp != act {
			sb.WriteString(fmt.Sprintf("line %d:\n  expected: %q\n  got:      %q\n", i+1, exp, act))
		}
	}
	return sb.String()
}

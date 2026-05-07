package judge

import (
	"os"
	"strings"

	"judge/internal/sandbox"
)

func JudgeSubmission(code, language, questionID string) (string, string) {
	inputPath := "Questions/" + questionID + "_in.txt"
	outputPath := "Questions/" + questionID + "_out.txt"

	inputData, err := os.ReadFile(inputPath)
	if err != nil {
		return "SYSTEM_ERROR", err.Error()
	}

	expectedOutput, err := os.ReadFile(outputPath)
	if err != nil {
		return "SYSTEM_ERROR", err.Error()
	}

	out, runtimeErr := sandbox.RunCode(
		code,
		language,
		string(inputData),
	)

	if runtimeErr == "TLE" {
		return "TLE", ""
	}

	if runtimeErr != "" {
		return "RUNTIME_ERROR", runtimeErr
	}

	if strings.TrimSpace(out) != strings.TrimSpace(string(expectedOutput)) {
		return "WRONG_ANSWER", "Output mismatch"
	}

	return "ACCEPTED", "All test cases passed"
}

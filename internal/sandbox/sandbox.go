package sandbox

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const (
	MemoryLimitMB   = 128
	TimeLimitSecond = 2
	SandboxImage    = "cpp-py-sandbox"
)

func RunCode(code, language, inputData string) (string, string) {
	tmpdir, err := os.MkdirTemp("", "judge-*")
	if err != nil {
		return "", err.Error()
	}
	defer os.RemoveAll(tmpdir)

	var sourceName string

	switch language {
	case "python3":
		sourceName = "main.py"
	case "cpp17":
		sourceName = "main.cpp"
	default:
		return "", "Unsupported language"
	}

	sourcePath := filepath.Join(tmpdir, sourceName)
	inputPath := filepath.Join(tmpdir, "input.txt")

	err = os.WriteFile(sourcePath, []byte(code), 0644)
	if err != nil {
		return "", err.Error()
	}

	err = os.WriteFile(inputPath, []byte(inputData), 0644)
	if err != nil {
		return "", err.Error()
	}

	if language == "python3" {
		return runPython(tmpdir)
	}

	return runCPP(tmpdir)
}

func runPython(tmpdir string) (string, string) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(TimeLimitSecond+2)*time.Second,
	)
	defer cancel()

	cmd := exec.CommandContext(
		ctx,
		"docker",
		"run",
		"--rm",
		"--network",
		"none",
		"--cap-drop",
		"ALL",
		"--security-opt",
		"no-new-privileges",
		"--memory",
		"128m",
		"--cpus",
		"1",
		"--read-only",
		"--tmpfs",
		"/tmp:rw,noexec,nosuid,size=50m",
		"-u",
		"1000:1000",
		"-v",
		tmpdir+":/workspace:rw",
		"-w",
		"/workspace",
		SandboxImage,
		"sh",
		"-c",
		"timeout 2 python3 main.py < input.txt",
	)

	output, err := cmd.CombinedOutput()

	if ctx.Err() == context.DeadlineExceeded {
		return "", "TLE"
	}

	if err != nil {
		if strings.Contains(string(output), "124") {
			return "", "TLE"
		}
		return "", string(output)
	}

	return string(output), ""
}

func runCPP(tmpdir string) (string, string) {
	compileCtx, compileCancel := context.WithTimeout(
		context.Background(),
		30*time.Second,
	)
	defer compileCancel()

	compileCmd := exec.CommandContext(
		compileCtx,
		"docker",
		"run",
		"--rm",
		"--network",
		"none",
		"--cap-drop",
		"ALL",
		"--security-opt",
		"no-new-privileges",
		"--memory",
		"128m",
		"--cpus",
		"1",
		"-u",
		"1000:1000",
		"-v",
		tmpdir+":/workspace:rw",
		"-w",
		"/workspace",
		SandboxImage,
		"g++",
		"main.cpp",
		"-std=c++17",
		"-O2",
		"-o",
		"program",
	)

	compileOutput, err := compileCmd.CombinedOutput()
	if err != nil {
		return "", string(compileOutput)
	}

	runCtx, runCancel := context.WithTimeout(
		context.Background(),
		time.Duration(TimeLimitSecond+2)*time.Second,
	)
	defer runCancel()

	runCmd := exec.CommandContext(
		runCtx,
		"docker",
		"run",
		"--rm",
		"--network",
		"none",
		"--cap-drop",
		"ALL",
		"--security-opt",
		"no-new-privileges",
		"--memory",
		"128m",
		"--cpus",
		"1",
		"--read-only",
		"--tmpfs",
		"/tmp:rw,noexec,nosuid,size=50m",
		"-u",
		"1000:1000",
		"-v",
		tmpdir+":/workspace:rw",
		"-w",
		"/workspace",
		SandboxImage,
		"sh",
		"-c",
		"timeout 2 ./program < input.txt",
	)

	output, err := runCmd.CombinedOutput()

	if runCtx.Err() == context.DeadlineExceeded {
		return "", "TLE"
	}

	if err != nil {
		return "", string(output)
	}

	return string(output), ""
}

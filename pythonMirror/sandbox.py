import os
import subprocess
import tempfile

from config import (
    MEMORY_LIMIT_MB,
    TIME_LIMIT_SECONDS,
)

PYTHON_IMAGE = "python:3.10-alpine"
# CPP_IMAGE = "alpine:3.18"
CPP_IMAGE = "cpp17-alpine"


def run_code(code: str, language: str, input_data: str):
    with tempfile.TemporaryDirectory() as tmpdir:
        if language == "python3":
            source_name = "main.py"

        elif language == "cpp17":
            source_name = "main.cpp"

        else:
            return "", "Unsupported language"

        source_path = os.path.join(tmpdir, source_name)

        with open(source_path, "w") as f:
            f.write(code)

        input_path = os.path.join(tmpdir, "input.txt")

        with open(input_path, "w") as f:
            f.write(input_data)

        try:
            if language == "python3":
                cmd = [
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
                    f"{MEMORY_LIMIT_MB}m",
                    "--cpus",
                    "1",
                    "--read-only",
                    "--tmpfs",
                    "/tmp",
                    "-v",
                    f"{tmpdir}:/workspace",
                    "-w",
                    "/workspace",
                    PYTHON_IMAGE,
                    "sh",
                    "-c",
                    f"timeout {TIME_LIMIT_SECONDS} python3 main.py",
                ]

                result = subprocess.run(
                    cmd,
                    input=input_data,
                    capture_output=True,
                    text=True,
                    timeout=TIME_LIMIT_SECONDS + 2,
                )

            elif language == "cpp17":
                # compile_cmd = [
                #     "docker",
                #     "run",
                #     "--rm",
                #     "--network",
                #     "none",
                #     "--cap-drop",
                #     "ALL",
                #     "--security-opt",
                #     "no-new-privileges",
                #     "--memory",
                #     f"{MEMORY_LIMIT_MB}m",
                #     "--cpus",
                #     "1",
                #     "-v",
                #     f"{tmpdir}:/workspace",
                #     "-w",
                #     "/workspace",
                #     CPP_IMAGE,
                #     "sh",
                #     "-c",
                #     "apk add --no-cache g++ musl-dev && "
                #     "g++ main.cpp -std=c++17 -O2 -o program",
                # ]
                compile_cmd = [
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
                    f"{MEMORY_LIMIT_MB}m",
                    "--cpus",
                    "1",
                    "-v",
                    f"{tmpdir}:/workspace",
                    "-w",
                    "/workspace",
                    CPP_IMAGE,
                    "g++",
                    "main.cpp",
                    "-std=c++17",
                    "-O2",
                    "-o",
                    "program",
                ]

                compile_result = subprocess.run(
                    compile_cmd,
                    capture_output=True,
                    text=True,
                    timeout=60,
                )

                if compile_result.returncode != 0:
                    return "", compile_result.stderr

                run_cmd = [
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
                    f"{MEMORY_LIMIT_MB}m",
                    "--cpus",
                    "1",
                    "--read-only",
                    "--tmpfs",
                    "/tmp",
                    "-v",
                    f"{tmpdir}:/workspace",
                    "-w",
                    "/workspace",
                    CPP_IMAGE,
                    "sh",
                    "-c",
                    f"timeout {TIME_LIMIT_SECONDS} ./program",
                ]

                result = subprocess.run(
                    run_cmd,
                    input=input_data,
                    capture_output=True,
                    text=True,
                    timeout=TIME_LIMIT_SECONDS + 10,
                )

            if result.returncode == 124:
                return "", "TLE"

            return result.stdout, result.stderr

        except subprocess.TimeoutExpired:
            return "", "TLE"

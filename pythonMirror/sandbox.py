import os
import subprocess
import tempfile

from config import MEMORY_LIMIT_MB, SANDBOX_IMAGE, TIME_LIMIT_SECONDS


def run_code(code: str, language: str, input_data: str):
    with tempfile.TemporaryDirectory() as tmpdir:
        if language == "python3":
            source_name = "main.py"
        elif language == "cpp17":
            source_name = "main.cpp"
        else:
            return "", "Unsupported language"

        source_path = os.path.join(tmpdir, source_name)
        input_path = os.path.join(tmpdir, "input.txt")

        with open(source_path, "w") as f:
            f.write(code)
        with open(input_path, "w") as f:
            f.write(input_data)

        os.chmod(source_path, 0o644)
        os.chmod(input_path, 0o644)

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
                    "/tmp:rw,noexec,nosuid,size=50m",
                    "-u",
                    "1000:1000",
                    "-v",
                    f"{tmpdir}:/workspace:rw",
                    "-w",
                    "/workspace",
                    SANDBOX_IMAGE,
                    "sh",
                    "-c",
                    f"timeout {TIME_LIMIT_SECONDS} python3 main.py < input.txt",
                ]

                result = subprocess.run(
                    cmd,
                    capture_output=True,
                    text=True,
                    timeout=TIME_LIMIT_SECONDS + 2,
                )

            elif language == "cpp17":
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
                    "-u",
                    "1000:1000",
                    "-v",
                    f"{tmpdir}:/workspace:rw",
                    "-w",
                    "/workspace",
                    SANDBOX_IMAGE,
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
                    timeout=30,
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
                    "/tmp:rw,noexec,nosuid,size=50m",
                    "-u",
                    "1000:1000",
                    "-v",
                    f"{tmpdir}:/workspace:rw",
                    "-w",
                    "/workspace",
                    SANDBOX_IMAGE,
                    "sh",
                    "-c",
                    f"timeout {TIME_LIMIT_SECONDS} ./program < input.txt",
                ]

                result = subprocess.run(
                    run_cmd,
                    capture_output=True,
                    text=True,
                    timeout=TIME_LIMIT_SECONDS + 5,
                )

            if result.returncode == 124:
                return "", "TLE"

            stderr = result.stderr
            if stderr and (
                "warning" not in stderr.lower() or "error" in stderr.lower()
            ):
                return result.stdout, stderr
            return result.stdout, ""

        except subprocess.TimeoutExpired:
            return "", "TLE"
        except Exception as e:
            return "", f"Sandbox error: {str(e)}"

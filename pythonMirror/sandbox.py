import subprocess
import tempfile
import os

from config import TIME_LIMIT_SECONDS


def run_code(code: str, language: str, input_data: str):
    if language == "python3":
        suffix = ".py"
    elif language == "cpp17":
        suffix = ".cpp"
    else:
        return "", "Unsupported language"

    with tempfile.TemporaryDirectory() as tmpdir:
        source_path = os.path.join(tmpdir, f"main{suffix}")

        with open(source_path, "w") as f:
            f.write(code)

        try:
            if language == "python3":
                result = subprocess.run(
                    ["python3", source_path],
                    input=input_data.encode(),
                    capture_output=True,
                    timeout=TIME_LIMIT_SECONDS
                )

            elif language == "cpp17":
                binary_path = os.path.join(tmpdir, "program")

                compile_result = subprocess.run(
                    [
                        "g++",
                        source_path,
                        "-std=c++17",
                        "-O2",
                        "-o",
                        binary_path
                    ],
                    capture_output=True
                )

                if compile_result.returncode != 0:
                    return "", compile_result.stderr.decode()

                result = subprocess.run(
                    [binary_path],
                    input=input_data.encode(),
                    capture_output=True,
                    timeout=TIME_LIMIT_SECONDS
                )

            return (
                result.stdout.decode(),
                result.stderr.decode()
            )

        except subprocess.TimeoutExpired:
            return "", "TLE"

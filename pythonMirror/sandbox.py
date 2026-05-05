import subprocess
import tempfile
import os
from config import TIME_LIMIT_SECONDS


def run_code(code: str, input_data: str):
    with tempfile.NamedTemporaryFile(delete=False, suffix=".py") as f:
        f.write(code.encode())
        file_path = f.name

    try:
        result = subprocess.run(
            ["python3", file_path],
            input=input_data.encode(),
            capture_output=True,
            timeout=TIME_LIMIT_SECONDS
        )
        return result.stdout.decode(), result.stderr.decode()

    except subprocess.TimeoutExpired:
        return "", "TLE"

    finally:
        os.remove(file_path)
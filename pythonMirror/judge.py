from sandbox import run_code

TEST_CASES = [
    ("2 3\n", "5\n"),
    ("10 20\n", "30\n"),
]


def judge_submission(code: str):
    for inp, expected in TEST_CASES:
        out, err = run_code(code, inp)

        if err == "TLE":
            return "TLE", ""

        if err:
            return "RUNTIME_ERROR", err

        if out.strip() != expected.strip():
            return "WRONG_ANSWER", f"Expected {expected}, got {out}"

    return "ACCEPTED", "All test cases passed"
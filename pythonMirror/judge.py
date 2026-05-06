from sandbox import run_code

def judge_submission(code, language, question_id):
    input_path = f"Questions/{question_id}_in.txt"
    output_path = f"Questions/{question_id}_out.txt"

    with open(input_path) as f:
        inp = f.read()

    with open(output_path) as f:
        expected = f.read()

    out, err = run_code(
        code,
        language,
        inp
    )

    if err == "TLE":
        return "TLE", ""

    if err:
        return "RUNTIME_ERROR", err

    if out.strip() != expected.strip():
        return (
            "WRONG_ANSWER",
            f"Expected {expected}, got {out}"
        )

    return "ACCEPTED", "All test cases passed"

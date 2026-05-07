from db import update_submission
from judge import judge_submission
import time

# replace with a proper queue later
QUEUE = []


def add_job(submission_id, code, language, question_id):
    QUEUE.append(
        (
            submission_id,
            code,
            language,
            question_id
        )
    )


def worker_loop():
    print("Worker started...")
    while True:
        if QUEUE:
            sub_id, code, language, question_id = QUEUE.pop(0)

            status, output = judge_submission(
                code,
                language,
                question_id
            )

            update_submission(sub_id, status, output)

            print(f"Processed submission {sub_id}: {status}")

        time.sleep(1)

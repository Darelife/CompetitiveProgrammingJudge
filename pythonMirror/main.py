"""
Competitive Programming Judge

1. API Endpoint for submitting code
2. Queues
3. Worker (async job processor)
4. Storage (DB for submissions + results)
5. Sandbox (Docker for execution)

log every step:
    1. Received submission
    2. Validation successful
    3. Storaged in DB
    4. Added to queue
    5. Worker <x> picked up job
    6. Execution Started
    7. Failures (if any)
    8. Execution Completed
    9. Execution Time
    10. Results stored in DB
    11. Results returned to client


submit -> store -> queue -> worker -> store results

Sandbox:
    - Time limits
    - Memory limits
    - Security isolation
    - Specific CPU speed (idk how to do this)
    - Worker crash -> job retry
"""

from typing import Literal, Union
from fastapi import FastAPI
from pydantic import BaseModel
import threading

from db import create_submission, get_submission
from worker import add_job, worker_loop

app = FastAPI()

class Submission(BaseModel):
    code: str
    language: Literal["python3", "cpp17"]
    question_id: str

@app.on_event("startup")
def start_worker():
    thread = threading.Thread(target=worker_loop, daemon=True)
    thread.start()


@app.post("/submit")
def submit_code(sub: Submission):
    sub_id = create_submission(
        sub.code,
        sub.language,
        sub.question_id
    )

    add_job(
        sub_id,
        sub.code,
        sub.language,
        sub.question_id
    )

    return {"submission_id": sub_id}


@app.get("/result/{submission_id}")
def get_result(submission_id: int):
    row = get_submission(submission_id)
    if not row:
        return {"error": "Not found"}

    return {
        "id": row[0],
        "language": row[2],
        "question_id": row[3],
        "status": row[4],
        "output": row[5]
    }

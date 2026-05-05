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

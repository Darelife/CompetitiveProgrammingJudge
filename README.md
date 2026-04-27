# CompetitiveProgrammingJudge

Learning how to use golang as a backend server + production level backend basics (caching, ratelimiting, batch processing, reverse proxy, databases, queues, and containerization)

## Todo

I'm using LLMs btw, trying to understand how everything works. I myself don't know (or have forgotten) about a lot of the imp things. So, i'll go and revise them

- [x] Revise [GoLang](GoLang.md)
- [ ] Revise Redis
- [ ] Revise Docker
- [ ] Learn how to use Nginx (Reverse Proxies)
- [ ] Learn how to use RabbitMQ
- [ ] Graceful shutdown & signal handling
- [ ] Rate limiting (token bucket, leaky bucket)
- [ ] Caching strategies
- [ ] External API failures
- [ ] GRPC

## Project Goal

Building a **competitive programming judge system**:
- Accept code submissions
- Run them securely in isolated environments
- Evaluate against test cases
- Return verdicts (AC / WA / TLE / RE)

Things to ensure:
- nginx API gateway (reverse proxy, and acts as a ratelimiter too)
- Go backend saves the metadata of the submission in local postgres, and pushes the job to a queue (perhaps rabbitmq)
- The judge instance keeps pulling jobs from the queue, sets up proper constraints, and security stuff (sandbox)
- Cache for tcs maybe?

## File Structure

- `cmd/` — Entry point for the API server
- `internal/` — Judging logic, transport handlers, and database layer
- `data/` — Test case input and expected output files

---

## Current Pipeline



```
┌─────────────────────────────────────────────────────────────────┐
│                        run.sh                                   │
│  1. go build ./cmd/server  →  /tmp/judge-server                 │
│  2. go build ./cmd/client  →  /tmp/judge-client                 │
│  3. Start server in background                                  │
│  4. Run client (makes 2 submissions)                            │
│  5. Kill server                                                 │
└────────────────┬────────────────────────────────────────────────┘
                 │
                 ▼
┌─────────────────────────────────────────────────────────────────┐
│  cmd/server/main.go   (HTTP server, :8080)                      │
│                                                                 │
│   Registers:  POST /submit  →  transport.HandleSubmission()     │
│   Listens on port 8080 and forwards every request to handler.   │
└────────────────┬────────────────────────────────────────────────┘
                 │  POST /submit  { "code": "<C++ source>" }
                 │◀──────────────────────────────────────────────
                 │                                cmd/client/main.go
                 │   (sends two submissions: one correct, one wrong)
                 ▼
┌─────────────────────────────────────────────────────────────────┐
│  internal/transport/handlers.go   (HTTP layer)                  │
│                                                                 │
│  [1] Decode JSON body                                           │
│        └─ bad/missing code  ──▶  400 Bad Request                │
│                                                                 │
│  [2] Delegate to judge.Evaluate(req.Code)                       │
└────────────────┬────────────────────────────────────────────────┘
                 ▼
┌─────────────────────────────────────────────────────────────────┐
│  internal/judge/judge.go   (the judge core)                     │
│                                                                 │
│  [3] Write req.Code to  /tmp/judge-XXXX/solution.cpp            │
│                                                                 │
│  [4] Compile                                                    │
│        g++ -o /tmp/judge-XXXX/solution  solution.cpp            │
│        └─ compile error  ──▶  422  +  compiler stderr           │
│                                                                 │
│  [5] Run                                                        │
│        ./solution  <  data/input.txt                            │
│        └─ runtime error  ──▶  422  +  error message             │
│                                                                 │
│  [6] Compare                                                    │
│        actual   ← stdout of ./solution  (whitespace trimmed)    │
│        expected ← data/expected_output.txt  (trimmed)           │
│                                                                 │
│  [7] Verdict                                                    │
│        actual == expected  ──▶  { verdict: "AC" }               │
│        actual != expected  ──▶  { verdict: "WA", diff: "..." }  │
└────────────────┬────────────────────────────────────────────────┘
                 │  returns models.SubmissionResponse
                 ▼
┌─────────────────────────────────────────────────────────────────┐
│  internal/transport/handlers.go   (HTTP layer)                  │
│                                                                 │
│  [8] Send JSON response                                         │
└────────────────┬────────────────────────────────────────────────┘
                 │  JSON response
                 ▼
┌─────────────────────────────────────────────────────────────────┐
│  internal/models/submission.go   (shared wire types)            │
│                                                                 │
│  SubmissionRequest   { code: string }                           │
│  SubmissionResponse  { verdict, output, diff? }                 │
└─────────────────────────────────────────────────────────────────┘
```

---

## File Map

| File | Role |
|---|---|
| `run.sh` | Builds everything, starts server, runs client, tears down |
| `cmd/server/main.go` | HTTP server entry point — registers routes, starts listener |
| `cmd/client/main.go` | Demo client — submits two C++ programs and prints verdicts |
| `internal/transport/handlers.go` | HTTP layer: parses request, delegates to judge, sends response |
| `internal/judge/judge.go` | Core judge logic: compile → run → compare |
| `internal/models/submission.go` | Shared request/response types (wire format) |
| `data/input.txt` | Test input fed to every submission (`1 2 3 4 5`) |
| `data/expected_output.txt` | Expected output compared against actual (`1 2 3 4 5`) |

---

## Running

```bash
bash run.sh
```

Expected output:
```
=== Building server and client ===
=== Starting judge server ===
Judge server running on :8080
=== Running client submissions ===

========== Submission: Correct solution (echo) ==========
✅ Verdict: AC (Accepted)
   Output : 1 2 3 4 5

========== Submission: Wrong solution (reversed) ==========
❌ Verdict: WA (Wrong Answer)
   Output : 5 4 3 2 1
   Diff   :
line 1:
  expected: "1 2 3 4 5"
  got:      "5 4 3 2 1"

=== Done. Stopping server. ===
```

---

## Verdict Reference

| Verdict | Meaning |
|---|---|
| `AC` | Accepted — output matches expected exactly |
| `WA` | Wrong Answer — output differs; `diff` field shows what's wrong |
| HTTP 422 | Compile error or runtime error (details in `error` field) |
| HTTP 400 | Bad request — missing `code` field |
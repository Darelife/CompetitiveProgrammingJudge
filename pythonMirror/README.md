# Running 

```bash
uvicorn main:app --host 0.0.0.0 --port 4000 --reload
```

# Curl Commands

```bash
curl -X POST "http://127.0.0.1:8000/submit" \
  -H "Content-Type: application/json" \
  -d '{
    "code": "print(sum(map(int, input().split())))"
  }'

curl -X POST "http://127.0.0.1:4000/submit" \
  -H "Content-Type: application/json" \
  -d '{
    "code": "print(sum(map(int, input().split())))",
    "language": "python3",
    "question_id": "a"
  }'

curl -X POST "http://127.0.0.1:4000/submit" \
  -H "Content-Type: application/json" \
  -d "$(jq -n \
    --rawfile code solution.cpp \
    '{
      code: $code,
      language: "cpp17",
      question_id: "a"
    }')"

curl -X POST "http://127.0.0.1:4000/submit" \
  -H "Content-Type: application/json" \
  -d "$(jq -n \
    --rawfile code solution.py \
    '{
      code: $code,
      language: "python3",
      question_id: "a"
    }')"
```

```bash
curl "http://127.0.0.1:4000/result/1"
```

<!--```
docker build -t cpp17-alpine .
```-->

```
docker build -f Dockerfile.sandbox -t cpp-py-sandbox .
```

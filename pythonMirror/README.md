# Curl Commands

```bash
curl -X POST "http://127.0.0.1:8000/submit" \
  -H "Content-Type: application/json" \
  -d '{
    "code": "print(sum(map(int, input().split())))"
  }'
```

```bash
curl "http://127.0.0.1:8000/result/1"
```